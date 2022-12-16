package handler

import (
	"encoding/json"
	"net/http"
	"path"
	"restapi/src/common/dto"
	"restapi/src/domain/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type ITagHandler interface {
	GetTags(http.ResponseWriter, *http.Request, httprouter.Params)
	CreateTag(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateTag(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteTag(http.ResponseWriter, *http.Request, httprouter.Params)
}

type TagHandler struct {
	service.ITagService
}

func NewTagHandler(srv service.ITagService) ITagHandler {
	return &TagHandler{srv}
}

func (h *TagHandler) GetTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tags, err := h.ITagService.GetTags()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var tag dto.Tag
	if err := json.Unmarshal(body, &tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTag, err := h.ITagService.CreateTag(&tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(createdTag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) UpdateTag(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var tag dto.Tag
	if err := json.Unmarshal(body, &tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTag, err := h.ITagService.UpdateTag(&tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedTag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) DeleteTag(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	tagId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.ITagService.DeleteTag(tagId); err == gorm.ErrRecordNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
