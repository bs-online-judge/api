package api_test

import (
  "net/http"
  "net/http/httptest"
  
  "github.com/bs-online-judge/api/api"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Middlewares", func() {
  var count int
  var w http.ResponseWriter
  var r *http.Request
  
  BeforeEach(func() {
    count = 0
    w = httptest.NewRecorder()
    r = httptest.NewRequest("GET", "/", nil)
  })
  
  trueFunc := func(*api.App) bool {
    count++
    return true
  }
  
  falseFunc := func(*api.App) bool {
    return false
  }
  
  It("should stop after first step of one", func() {
    api.Chain(falseFunc)(w, r)
    Expect(count).To(Equal(0))
  })
  
  It("should stop after all steps of one", func() {
    api.Chain(trueFunc)(w, r)
    Expect(count).To(Equal(1))
  })
  
  It("should stop after first step of two", func() {
    api.Chain(falseFunc, trueFunc)(w, r)
    Expect(count).To(Equal(0))
  })
  
  It("should stop after second step of two", func() {
    api.Chain(trueFunc, falseFunc)(w, r)
    Expect(count).To(Equal(1))
  })
  
  It("should stop after all steps of two", func() {
    api.Chain(trueFunc, trueFunc)(w, r)
    Expect(count).To(Equal(2))
  })
  
  It("should stop after first step of three", func() {
    api.Chain(falseFunc, trueFunc, trueFunc)(w, r)
    Expect(count).To(Equal(0))
  })
  
  It("should stop after second step of three", func() {
    api.Chain(trueFunc, falseFunc, trueFunc)(w, r)
    Expect(count).To(Equal(1))
  })
  
  It("should stop after third step of three", func() {
    api.Chain(trueFunc, trueFunc, falseFunc)(w, r)
    Expect(count).To(Equal(2))
  })
  
  It("should stop after all steps of three", func() {
    api.Chain(trueFunc, trueFunc, trueFunc)(w, r)
    Expect(count).To(Equal(3))
  })
})
