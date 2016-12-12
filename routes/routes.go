package routes

import (
	"github.com/dictyBase/go-middlewares/middlewares/chain"
	"github.com/dictyBase/go-middlewares/middlewares/pagination"
	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/dictyBase/modware/resources"
)

func AddPublication(rs resources.Resource, mwChain chain.Chain, r *router.Wrapper) {
	r.Get("/publications/:id", mwChain.ThenFunc(rs.Get))
	r.Get("/publications", mwChain.Append(pagination.MiddlewareFn).ThenFunc(rs.GetAll))
	r.Post("/publications", mwChain.ThenFunc(rs.Create))
	r.Patch("/publications/:id", mwChain.ThenFunc(rs.Update))
	r.Delete("/publication/:id", mwChain.ThenFunc(rs.Delete))
}

func AddAuthor(rs resources.Resource, mwChain chain.Chain, r *router.Wrapper) {

}
