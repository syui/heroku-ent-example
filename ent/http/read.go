// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"t/ent"
	"t/ent/compartment"
	"t/ent/entry"
	"t/ent/fridge"
	"t/ent/item"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Read fetches the ent.Compartment identified by a given url-parameter from the
// database and renders it to the client.
func (h *CompartmentHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the Compartment
	q := h.client.Compartment.Query().Where(compartment.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read compartment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("compartment rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment3324871446View(e), w)
}

// Read fetches the ent.Entry identified by a given url-parameter from the
// database and renders it to the client.
func (h *EntryHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	var err error
	id := chi.URLParam(r, "id")
	// Create the query to fetch the Entry
	q := h.client.Entry.Query().Where(entry.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.String("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read entry", zap.Error(err), zap.String("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("entry rendered", zap.String("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewEntry2675665849View(e), w)
}

// Read fetches the ent.Fridge identified by a given url-parameter from the
// database and renders it to the client.
func (h *FridgeHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the Fridge
	q := h.client.Fridge.Query().Where(fridge.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read fridge", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("fridge rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2211356377View(e), w)
}

// Read fetches the ent.Item identified by a given url-parameter from the
// database and renders it to the client.
func (h *ItemHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the Item
	q := h.client.Item.Query().Where(item.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read item", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("item rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewItem1548468123View(e), w)
}
