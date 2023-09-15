package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"blacheapi/http/graphql/model"
	"blacheapi/http/rest/interceptor"
	"blacheapi/logger"
	"blacheapi/logic"
	"blacheapi/primer/constant"
	"blacheapi/primer/enum"
	"blacheapi/primer/function"
	"blacheapi/primer/gql"
	"blacheapi/primer/typing"
	"blacheapi/repository"
	"context"
	"errors"
	"fmt"
)

// Webhook is the resolver for the webhook field.
func (r *mutationResolver) Webhook(ctx context.Context, input model.WebhookInput) (model.RespondWithTransaction, error) {
	// validation
	// validation := validate(schema, input)
	// if !validation.Status {
	// nil, gql.MakeSubgraphError(validation.Message, http.StatusBadRequest)
	// }

	act := repository.Activity{
		ID:            function.GenerateUUID(),
		Resolver:      "mutationResolver.Webhook",
		Payload:       function.Stringify(input),
		Description:   "made an action to verify a transaction",
		TransactionID: "",
		SavingsID:     "",
		Role:          "",
		Error:         "",
		Status:        "",
	}
	act.Date()

	logger.GetLogger().Debug(fmt.Sprintf(`START :: [%v] :: mutationResolver.Webhook with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	tx, err := logic.Webhook(input, false)
	if err != nil {
		if errors.As(err, &gql.Error{}) {
			return gql.MakeSubgraphError(err.(gql.Error).Message, err.(gql.Error).Status, fmt.Sprintf(`[%v] :: {message}`, ctx.Value(typing.CtxTraceKey{})), function.Stringify(act)), nil
		}

		return gql.MakeSubgraphError(`Something went wrong while trying to verify your transaction! Please try again.`, constant.CodeISE, fmt.Sprintf(`[%v] :: {message} :: %s`, ctx.Value(typing.CtxTraceKey{}), err.Error()), function.Stringify(act)), nil
	}

	logger.GetLogger().Debug(fmt.Sprintf(`END :: [%v] :: mutationResolver.Webhook with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	// log activity
	act.Status = enum.Success
	act.TransactionID = tx.ID
	act.SavingsID = tx.SavingsID
	if err := act.Create(); err != nil {
		logger.GetLogger().Debug(`ACTIVITY LOG :: SAVE ERROR :: ` + err.Error())
	}

	var ot model.Transaction
	function.Parse(tx, &ot)
	return &model.ResponseWithTransaction{
		Message: "Transaction verified.",
		Data:    &ot,
	}, nil
}

// Transaction is the resolver for the transaction field.
func (r *queryResolver) Transaction(ctx context.Context, input model.TransactionFilterInput) (model.RespondWithTransaction, error) {
	logger.GetLogger().Debug(fmt.Sprintf(`START :: [%v] :: queryResolver.Transaction with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	act := repository.Activity{
		ID:            function.GenerateUUID(),
		Resolver:      "queryResolver.Transaction",
		Payload:       function.Stringify(input),
		Description:   "made an action to retrieve a transaction",
		TransactionID: "",
		Role:          "",
		Error:         "",
		Status:        "",
	}
	if input.TransactionID != nil {
		act.TransactionID = *input.TransactionID
	}
	act.Date()

	// access control
	auth, _, err := interceptor.Authorize(ctx)
	// we only expect a gql.Error here
	if err != nil && errors.As(err, &gql.Error{}) {
		return nil, gql.MakeGraphQLError(ctx, err.(gql.Error).Message, err.(gql.Error).Status, fmt.Sprintf(`[%v] :: {message}`, ctx.Value(typing.CtxTraceKey{})), function.Stringify(act))
	}

	act.Role = auth.Role
	act.By = &auth.ID

	// validation
	// validation := validate(schema, input)
	// if !validation.Status {
	// nil, gql.MakeSubgraphError(validation.Message, http.StatusBadRequest)
	// }

	// get transaction
	transaction, err := logic.Transaction(input)
	if err != nil {
		if errors.As(err, &gql.Error{}) {
			return gql.MakeSubgraphError(err.(gql.Error).Message, err.(gql.Error).Status, fmt.Sprintf(`[%v] :: {message}`, ctx.Value(typing.CtxTraceKey{})), function.Stringify(act)), nil
		}

		return gql.MakeSubgraphError(`Something went wrong while retrieving transaction! Please try again.`, constant.CodeISE, fmt.Sprintf(`[%v] :: {message} :: %s`, ctx.Value(typing.CtxTraceKey{}), err.Error()), function.Stringify(act)), nil
	}

	logger.GetLogger().Debug(fmt.Sprintf(`END :: [%v] :: queryResolver.Transaction with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	// log activity
	act.Status = enum.Success
	if err := act.Create(); err != nil {
		logger.GetLogger().Debug(`ACTIVITY LOG :: SAVE ERROR :: ` + err.Error())
	}

	var ot model.Transaction
	function.Parse(transaction, &ot)
	return &model.ResponseWithTransaction{
		Message: "Transaction retrieved.",
		Data:    &ot,
	}, nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, input *model.TransactionFilterInput) (model.RespondWithTransactions, error) {
	logger.GetLogger().Debug(fmt.Sprintf(`START :: [%v] :: queryResolver.Transactions with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	act := repository.Activity{
		ID:            function.GenerateUUID(),
		Resolver:      "queryResolver.Transactions",
		Payload:       function.Stringify(input),
		Description:   "made an action to retrieve multiple transactions",
		TransactionID: "",
		Error:         "",
		Status:        "",
	}
	if input != nil && input.TransactionID != nil {
		act.TransactionID = *input.TransactionID
	}
	act.Date()

	// access control
	auth, _, err := interceptor.Authorize(ctx)
	// we only expect a gql.Error here
	if err != nil && errors.As(err, &gql.Error{}) {
		return nil, gql.MakeGraphQLError(ctx, err.(gql.Error).Message, err.(gql.Error).Status, fmt.Sprintf(`[%v] :: {message}`, ctx.Value(typing.CtxTraceKey{})), function.Stringify(act))
	}

	act.Role = auth.Role
	act.By = &auth.ID

	// validation
	// validation := validate(schema, input)
	// if !validation.Status {
	// nil, gql.MakeSubgraphError(validation.Message, http.StatusBadRequest)
	// }

	// get transactions
	// @TODO: get customer from ctx and replace for ""
	transactions, pagination, err := logic.Transactions(input, repository.Customer{OrgID: auth.OrgID, Role: auth.Role})
	if err != nil {
		if errors.As(err, &gql.Error{}) {
			return gql.MakeSubgraphError(err.(gql.Error).Message, err.(gql.Error).Status, fmt.Sprintf(`[%v] :: {message}`, ctx.Value(typing.CtxTraceKey{})), function.Stringify(act)), nil
		}

		return gql.MakeSubgraphError(`Something went wrong while retrieving transactions! Please try again.`, constant.CodeISE, fmt.Sprintf(`[%v] :: {message} :: %s`, ctx.Value(typing.CtxTraceKey{}), err.Error()), function.Stringify(act)), nil
	}

	logger.GetLogger().Debug(fmt.Sprintf(`END :: [%v] :: queryResolver.Transactions with input: %+v`, ctx.Value(typing.CtxTraceKey{}), function.Stringify(input)))

	// log activity
	act.Status = enum.Success
	if err := act.Create(); err != nil {
		logger.GetLogger().Debug(`ACTIVITY LOG :: SAVE ERROR :: ` + err.Error())
	}

	var ots []*model.Transaction
	function.Parse(transactions, &ots)
	var ot *model.Pagination
	function.Parse(pagination, &ot)
	return &model.ResponseWithTransactions{
		Message:    "Transactions retrieved.",
		Data:       ots,
		Pagination: ot,
	}, nil
}
