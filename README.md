# Leetcode

## Idea plugin leetcode-editor

### **Golang**

```bash
## CodeFileName
$!{question.frontendQuestionId}.$!velocityTool.camelCaseName(${question.titleSlug})_test

## CodeTemplate
package Leetcode

import (
    "log"
    "testing"
)

func Test$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T) {
    fn := $!velocityTool.smallCamelCaseName(${question.titleSlug})
    log.Println(fn())
}

${question.code}
```

### Java

```bash
## CodeFileName
$!{question.frontendQuestionId}.$!velocityTool.camelCaseName(${question.titleSlug})_test

## CodeTemplate

```

### C++

```bash
## CodeFileName
$!{question.frontendQuestionId}.$!velocityTool.camelCaseName(${question.titleSlug})_test

## CodeTemplate

```







