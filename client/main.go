package main

import (
	"api_grpc_protobuf/proto"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewCreateServerClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(c *gin.Context) {
		a, err := strconv.ParseInt(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "inval Param A",
			})
		}
		b, err := strconv.ParseInt(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "inval Param B",
			})
		}
		req := &proto.Request{A: a, B: b}
		res, err := client.Add(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("client.Add() error(%v)", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		}
	})

	g.GET("/mult/:a/:b", func(c *gin.Context) {
		a, err := strconv.ParseInt(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "inval Param A",
			})
		}
		b, err := strconv.ParseInt(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "inval Param B",
			})
		}
		req := &proto.Request{A: a, B: b}
		res, err := client.Multiply(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("client.Multiply() error(%v)", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
