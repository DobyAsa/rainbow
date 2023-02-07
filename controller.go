package main

import (
	"context"
	"log"
	"net/http"
	"rainbow/framework"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		ctx.JSON(http.StatusOK, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		log.Println(p)
		ctx.JSON(http.StatusInternalServerError, "panic")
	case <-finish:
		log.Println("finish")
	case <-durationCtx.Done():
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		ctx.JSON(http.StatusInternalServerError, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}
