//go:generate mockgen  -destination=./mocks/mocks.go -package=mocks github.com/blocknative/dreamboat/pkg/api Relay,Registrations

package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flashbots/go-boost-utils/types"
	"github.com/gorilla/mux"
	"github.com/lthibault/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/blocknative/dreamboat/pkg/relay"
	"github.com/blocknative/dreamboat/pkg/structs"
	"github.com/blocknative/dreamboat/pkg/validators"
)

// Router paths
const (
	// proposer endpoints
	PathStatus            = "/eth/v1/builder/status"
	PathRegisterValidator = "/eth/v1/builder/validators"
	PathGetHeader         = "/eth/v1/builder/header/{slot:[0-9]+}/{parent_hash:0x[a-fA-F0-9]+}/{pubkey:0x[a-fA-F0-9]+}"
	PathGetPayload        = "/eth/v1/builder/blinded_blocks"

	// builder endpoints
	PathGetValidators = "/relay/v1/builder/validators"
	PathSubmitBlock   = "/relay/v1/builder/blocks"

	// data api
	PathBuilderBlocksReceived     = "/relay/v1/data/bidtraces/builder_blocks_received"
	PathProposerPayloadsDelivered = "/relay/v1/data/bidtraces/proposer_payload_delivered"
	PathSpecificRegistration      = "/relay/v1/data/validator_registration"
)

const (
	DataLimit = 450
)

var (
	ErrParamNotFound = errors.New("not found")
)

type Relay interface {
	// Proposer APIs (builder spec https://github.com/ethereum/builder-specs)
	GetHeader(context.Context, *structs.MetricGroup, structs.HeaderRequest) (*types.GetHeaderResponse, error)
	GetPayload(context.Context, *structs.MetricGroup, *types.SignedBlindedBeaconBlock) (*types.GetPayloadResponse, error)

	// Builder APIs (relay spec https://flashbots.notion.site/Relay-API-Spec-5fb0819366954962bc02e81cb33840f5)
	SubmitBlock(context.Context, *structs.MetricGroup, *types.BuilderSubmitBlockRequest) error

	// Data APIs
	GetPayloadDelivered(context.Context, structs.PayloadTraceQuery) ([]structs.BidTraceExtended, error)
	GetBlockReceived(context.Context, structs.HeaderTraceQuery) ([]structs.BidTraceWithTimestamp, error)
}

type Registrations interface {
	// Proposer APIs (builder spec https://github.com/ethereum/builder-specs)
	RegisterValidator(context.Context, *structs.MetricGroup, []types.SignedValidatorRegistration) error
	// Data APIs
	Registration(context.Context, types.PublicKey) (types.SignedValidatorRegistration, error)
	// Builder APIs (relay spec https://flashbots.notion.site/Relay-API-Spec-5fb0819366954962bc02e81cb33840f5)
	GetValidators(*structs.MetricGroup) structs.BuilderGetValidatorsResponseEntrySlice
}

type API struct {
	l   log.Logger
	r   Relay
	reg Registrations

	m APIMetrics
}

func NewApi(l log.Logger, r Relay, reg Registrations) (a *API) {
	a = &API{l: l, r: r, reg: reg}
	a.initMetrics()
	return a
}

func (a *API) AttachToHandler(m *http.ServeMux) {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router), withAddons(a.l))

	// root returns 200 - nil
	router.HandleFunc("/", status)

	// proposer related
	router.HandleFunc(PathStatus, status).Methods(http.MethodGet)
	router.HandleFunc(PathRegisterValidator, handler(a.registerValidator)).Methods(http.MethodPost)
	router.HandleFunc(PathGetHeader, handler(a.getHeader)).Methods(http.MethodGet)
	router.HandleFunc(PathGetPayload, handler(a.getPayload)).Methods(http.MethodPost)

	// builder related
	router.HandleFunc(PathSubmitBlock, handler(a.submitBlock)).Methods(http.MethodPost)
	router.HandleFunc(PathGetValidators, handler(a.getValidators)).Methods(http.MethodGet)

	// data API related
	router.HandleFunc(PathProposerPayloadsDelivered, handler(a.proposerPayloadsDelivered)).Methods(http.MethodGet)
	router.HandleFunc(PathBuilderBlocksReceived, handler(a.builderBlocksReceived)).Methods(http.MethodGet)
	router.HandleFunc(PathSpecificRegistration, handler(a.specificRegistration)).Methods(http.MethodGet)

	router.Use(mux.CORSMethodMiddleware(router))
	m.Handle("/", router)
}

func handler(f func(http.ResponseWriter, *http.Request) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := f(w, r)

		if status == 0 || status == http.StatusOK {
			status = http.StatusOK
		}

		w.WriteHeader(status)

		if err != nil {
			_ = json.NewEncoder(w).Encode(jsonError{
				Code:    status,
				Message: err.Error(),
			})
		}
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// proposer related handlers
func (a *API) registerValidator(w http.ResponseWriter, r *http.Request) (status int, err error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("registerValidator"))
	defer timer.ObserveDuration()

	payload := []types.SignedValidatorRegistration{}
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		a.m.ApiReqCounter.WithLabelValues("registerValidator", "400", "input decoding").Inc()
		return http.StatusBadRequest, errors.New("invalid payload")
	}

	if payload != nil {
		a.m.ApiReqElCount.WithLabelValues("registerValidator", "payload").Observe(float64(len(payload)))
	}

	m := structs.NewMetricGroup(4)
	if err = a.reg.RegisterValidator(r.Context(), m, payload); err != nil {
		m.ObserveWithError(a.m.RelayTiming, unwrapError(err, "register validator unknown"))
		a.m.ApiReqCounter.WithLabelValues("registerValidator", "400", "register validator").Inc()
		a.l.With(log.F{
			"code":     400,
			"endpoint": "registerValidator",
			"type":     "single",
			"payload":  payload,
		}).WithError(err).Debug("failed registerValidator")
		return http.StatusBadRequest, err
	}

	m.Observe(a.m.RelayTiming)
	a.m.ApiReqCounter.WithLabelValues("registerValidator", "200", "").Inc()
	return http.StatusOK, err
}

func (a *API) getHeader(w http.ResponseWriter, r *http.Request) (int, error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("getHeader"))
	defer timer.ObserveDuration()

	req := ParseHeaderRequest(r)
	m := structs.NewMetricGroup(4)
	response, err := a.r.GetHeader(r.Context(), m, req)
	if err != nil {
		m.ObserveWithError(a.m.RelayTiming, unwrapError(err, "get header unknown"))
		a.m.ApiReqCounter.WithLabelValues("getHeader", "400", "get header").Inc()
		slot, _ := req.Slot()
		proposer, _ := req.Pubkey()
		a.l.With(log.F{
			"code":     400,
			"endpoint": "getHeader",
			"payload":  req,
			"slot":     slot,
			"proposer": proposer,
		}).WithError(err).Debug("failed getHeader")
		return http.StatusBadRequest, err
	}

	m.Observe(a.m.RelayTiming)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		a.l.WithError(err).WithField("path", r.URL.Path).Debug("failed to write response")
		a.m.ApiReqCounter.WithLabelValues("getHeader", "500", "response encode").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("getHeader", "200", "").Inc()
	return http.StatusOK, nil
}

func (a *API) getPayload(w http.ResponseWriter, r *http.Request) (int, error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("getPayload"))
	defer timer.ObserveDuration()

	var req types.SignedBlindedBeaconBlock
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.m.ApiReqCounter.WithLabelValues("getPayload", "400", "payload decode").Inc()
		return http.StatusBadRequest, errors.New("invalid payload")
	}

	m := structs.NewMetricGroup(4)
	payload, err := a.r.GetPayload(r.Context(), m, &req)
	if err != nil {
		m.ObserveWithError(a.m.RelayTiming, unwrapError(err, "get payload unknown"))
		a.m.ApiReqCounter.WithLabelValues("getPayload", "400", "get payload").Inc()
		a.l.With(log.F{
			"code":      400,
			"endpoint":  "getPayload",
			"payload":   req,
			"slot":      req.Message.Slot,
			"blockHash": req.Message.Body.Eth1Data.BlockHash,
		}).WithError(err).Debug("failed getPayload")
		return http.StatusBadRequest, err
	}

	m.Observe(a.m.RelayTiming)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		a.l.WithError(err).WithField("path", r.URL.Path).Debug("failed to write response")
		a.m.ApiReqCounter.WithLabelValues("getPayload", "500", "encode response").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("getPayload", "200", "").Inc()
	return http.StatusOK, nil
}

// builder related handlers
func (a *API) submitBlock(w http.ResponseWriter, r *http.Request) (int, error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("submitBlock"))
	defer timer.ObserveDuration()

	var req types.BuilderSubmitBlockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.m.ApiReqCounter.WithLabelValues("submitBlock", "400", "payload decode").Inc()
		return http.StatusBadRequest, errors.New("invalid payload")
	}

	if req.ExecutionPayload != nil && req.ExecutionPayload.Transactions != nil {
		a.m.ApiReqElCount.WithLabelValues("submitBlock", "transaction").Observe(float64(len(req.ExecutionPayload.Transactions)))
	}

	m := structs.NewMetricGroup(4)
	if err := a.r.SubmitBlock(r.Context(), m, &req); err != nil {
		m.ObserveWithError(a.m.RelayTiming, unwrapError(err, "submit block unknown"))
		if errors.Is(err, relay.ErrPayloadAlreadyDelivered) {
			a.m.ApiReqCounter.WithLabelValues("submitBlock", "400", "payload already delivered").Inc()
		} else {
			a.m.ApiReqCounter.WithLabelValues("submitBlock", "400", "block submission").Inc()
			a.l.With(log.F{
				"code":      400,
				"endpoint":  "submitBlock",
				"payload":   req,
				"slot":      req.Message.Slot,
				"blockHash": req.ExecutionPayload.BlockHash,
				"bidValue":  req.Message.Value,
				"proposer":  req.Message.ProposerPubkey,
				"builder":   req.Message.BuilderPubkey,
			}).WithError(err).Debug("failed block submission")
		}
		return http.StatusBadRequest, err
	}

	m.Observe(a.m.RelayTiming)
	a.m.ApiReqCounter.WithLabelValues("submitBlock", "200", "").Inc()
	return http.StatusOK, nil
}

func (a *API) getValidators(w http.ResponseWriter, r *http.Request) (int, error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("getValidators"))
	defer timer.ObserveDuration()

	m := structs.NewMetricGroup(4)
	vs := a.reg.GetValidators(m)
	if vs == nil {
		a.l.Trace("no registered validators for epoch")
		vs = structs.BuilderGetValidatorsResponseEntrySlice{}
		m.ObserveWithError(a.m.RelayTiming, fmt.Errorf("no validators"))
	}

	if vs != nil {
		m.Observe(a.m.RelayTiming)
		a.m.ApiReqElCount.WithLabelValues("getValidators", "validator").Observe(float64(len(vs)))
	}

	if err := json.NewEncoder(w).Encode(vs); err != nil {
		a.m.ApiReqCounter.WithLabelValues("getValidators", "500", "response encode").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("getValidators", "200", "").Inc()
	return http.StatusOK, nil
}

// data API related handlers
func (a *API) specificRegistration(w http.ResponseWriter, r *http.Request) (int, error) {
	timer := prometheus.NewTimer(a.m.ApiReqTiming.WithLabelValues("specificRegistration"))
	defer timer.ObserveDuration()

	pkStr := r.URL.Query().Get("pubkey")

	var pk types.PublicKey
	if err := pk.UnmarshalText([]byte(pkStr)); err != nil {
		a.m.ApiReqCounter.WithLabelValues("specificRegistration", "400", "unmarshaling pk").Inc()
		return http.StatusBadRequest, err
	}

	registration, err := a.reg.Registration(r.Context(), pk)
	if err != nil {
		a.m.ApiReqCounter.WithLabelValues("specificRegistration", "500", "registration").Inc()
		return http.StatusInternalServerError, err
	}

	if err := json.NewEncoder(w).Encode(registration); err != nil {
		a.m.ApiReqCounter.WithLabelValues("specificRegistration", "500", "encode response").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("specificRegistration", "200", "").Inc()
	return http.StatusOK, nil
}

func (a *API) proposerPayloadsDelivered(w http.ResponseWriter, r *http.Request) (int, error) {
	slot, err := specificSlot(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad slot").Inc()
		return http.StatusBadRequest, err
	}

	bh, err := blockHash(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad hash").Inc()
		return http.StatusBadRequest, err
	}

	bn, err := blockNumber(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad number").Inc()
		return http.StatusBadRequest, err
	}

	pk, err := publickKey(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad key").Inc()
		return http.StatusBadRequest, err
	}

	limit, err := limit(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad limit").Inc()
		return http.StatusBadRequest, err
	} else if errors.Is(err, ErrParamNotFound) {
		limit = DataLimit
	}

	cursor, err := cursor(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "400", "bad cursor").Inc()
		return http.StatusBadRequest, err
	}

	query := structs.PayloadTraceQuery{
		Slot:      slot,
		BlockHash: bh,
		BlockNum:  bn,
		Pubkey:    pk,
		Cursor:    cursor,
		Limit:     limit,
	}

	payloads, err := a.r.GetPayloadDelivered(r.Context(), query)
	if err != nil {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "500", "get payloads").Inc()
		return http.StatusInternalServerError, err
	}

	if payloads != nil {
		a.m.ApiReqElCount.WithLabelValues("proposerPayloadsDelivered", "payload").Observe(float64(len(payloads)))
	}

	if err := json.NewEncoder(w).Encode(payloads); err != nil {
		a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "500", "encode response").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("proposerPayloadsDelivered", "200", "").Inc()
	return http.StatusOK, nil
}

func (a *API) builderBlocksReceived(w http.ResponseWriter, r *http.Request) (int, error) {

	slot, err := specificSlot(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "400", "bad slot").Inc()
		return http.StatusBadRequest, err
	}

	bh, err := blockHash(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "400", "bad hash").Inc()
		return http.StatusBadRequest, err
	}

	bn, err := blockNumber(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "400", "bad number").Inc()
		return http.StatusBadRequest, err
	}

	limit, err := limit(r)
	if isInvalidParameter(err) {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "400", "bad limit").Inc()
		return http.StatusBadRequest, err
	} else if errors.Is(err, ErrParamNotFound) {
		limit = DataLimit
	}

	query := structs.HeaderTraceQuery{
		Slot:      slot,
		BlockHash: bh,
		BlockNum:  bn,
		Limit:     limit,
	}

	blocks, err := a.r.GetBlockReceived(r.Context(), query)
	if err != nil {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "500", "get block").Inc()
		return http.StatusInternalServerError, err
	}

	if blocks != nil {
		a.m.ApiReqElCount.WithLabelValues("builderBlocksReceived", "block").Observe(float64(len(blocks)))
	}

	if err := json.NewEncoder(w).Encode(blocks); err != nil {
		a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "500", "encode response").Inc()
		return http.StatusInternalServerError, err
	}

	a.m.ApiReqCounter.WithLabelValues("builderBlocksReceived", "200", "").Inc()
	return http.StatusOK, nil
}

func isInvalidParameter(err error) bool {
	return err != nil && !errors.Is(err, ErrParamNotFound)
}

func specificSlot(r *http.Request) (structs.Slot, error) {
	if slotStr := r.URL.Query().Get("slot"); slotStr != "" {
		slot, err := strconv.ParseUint(slotStr, 10, 64)
		if err != nil {
			return structs.Slot(0), err
		}
		return structs.Slot(slot), nil
	}
	return structs.Slot(0), ErrParamNotFound
}

func blockHash(r *http.Request) (types.Hash, error) {
	if bhStr := r.URL.Query().Get("block_hash"); bhStr != "" {
		var bh types.Hash
		if err := bh.UnmarshalText([]byte(bhStr)); err != nil {
			return bh, err
		}
		return bh, nil
	}
	return types.Hash{}, ErrParamNotFound
}

func publickKey(r *http.Request) (types.PublicKey, error) {
	if pkStr := r.URL.Query().Get("proposer_pubkey"); pkStr != "" {
		var pk types.PublicKey
		if err := pk.UnmarshalText([]byte(pkStr)); err != nil {
			return pk, err
		}
		return pk, nil
	}
	return types.PublicKey{}, ErrParamNotFound
}

func blockNumber(r *http.Request) (uint64, error) {
	if bnStr := r.URL.Query().Get("block_number"); bnStr != "" {
		return strconv.ParseUint(bnStr, 10, 64)
	}
	return 0, ErrParamNotFound
}

func limit(r *http.Request) (uint64, error) {
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err := strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			return 0, err
		} else if DataLimit < limit {
			return 0, fmt.Errorf("limit is higher than %d", DataLimit)
		}
		return limit, err
	}
	return 0, ErrParamNotFound
}

func cursor(r *http.Request) (uint64, error) {
	if cursorStr := r.URL.Query().Get("cursor"); cursorStr != "" {
		return strconv.ParseUint(cursorStr, 10, 64)
	}
	return 0, ErrParamNotFound
}

func ParseHeaderRequest(r *http.Request) structs.HeaderRequest {
	return mux.Vars(r)
}

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func unwrapError(err error, defaultMsg string) error {
	if errors.Is(err, relay.ErrUnknownValue) {
		return relay.ErrUnknownValue
	} else if errors.Is(err, relay.ErrPayloadAlreadyDelivered) {
		return relay.ErrPayloadAlreadyDelivered
	} else if errors.Is(err, relay.ErrNoPayloadFound) {
		return relay.ErrNoPayloadFound
	} else if errors.Is(err, relay.ErrMissingRequest) {
		return relay.ErrMissingRequest
	} else if errors.Is(err, relay.ErrMissingSecretKey) {
		return relay.ErrMissingSecretKey
	} else if errors.Is(err, relay.ErrNoBuilderBid) {
		return relay.ErrNoBuilderBid
	} else if errors.Is(err, relay.ErrOldSlot) {
		return relay.ErrOldSlot
	} else if errors.Is(err, relay.ErrBadHeader) {
		return relay.ErrBadHeader
	} else if errors.Is(err, relay.ErrInvalidSignature) {
		return relay.ErrInvalidSignature
	} else if errors.Is(err, relay.ErrStore) {
		return relay.ErrStore
	} else if errors.Is(err, relay.ErrMarshal) {
		return relay.ErrMarshal
	} else if errors.Is(err, relay.ErrInternal) {
		return relay.ErrInternal
	} else if errors.Is(err, relay.ErrUnknownValidator) {
		return relay.ErrUnknownValidator
	} else if errors.Is(err, relay.ErrVerification) {
		return relay.ErrVerification
	} else if errors.Is(err, relay.ErrInvalidTimestamp) {
		return relay.ErrInvalidTimestamp
	} else if errors.Is(err, relay.ErrInvalidSlot) {
		return relay.ErrInvalidSlot
	} else if errors.Is(err, relay.ErrEmptyBlock) {
		return relay.ErrEmptyBlock
	} else if errors.Is(err, validators.ErrInvalidSignature) {
		return validators.ErrInvalidSignature
	} else if errors.Is(err, validators.ErrUnknownValidator) {
		return validators.ErrUnknownValidator
	} else if errors.Is(err, validators.ErrInvalidTimestamp) {
		return validators.ErrInvalidTimestamp
	}

	return errors.New(defaultMsg)
}
