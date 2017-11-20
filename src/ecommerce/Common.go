package ecommerce

import (
  "bytes"
	"sort"
)

func encodeChar(ch string) string {
  switch ch {
    case "!":
      return "%21"
    case "#":
      return "%23"
    case "$":
      return "%24"
    case "&":
      return "%26"
    case "'":
      return "%27"
    case "(":
      return "%28"
    case ")":
      return "%29"
    case "*":
      return "%2A"
    case "+":
      return "%2B"
    case ",":
      return "%2C"
    case "/":
      return "%2F"
    case ":":
      return "%3A"
    case ";":
      return "%3B"
    case "=":
      return "%3D"
    case "?":
      return "%3F"
    case "@":
      return "%40"
    case "[":
      return "%5B"
    case "]":
      return "%5D"
    default:
      return ch
  }
}

func encodeUrl(text string) string {
  var buffer bytes.Buffer
  for _, r := range text {
      c := string(r)
      t := encodeChar(c)
      buffer.WriteString(t)
  }

  return buffer.String()
}

func generateParams(list map[string]string) (string, error) {
   var buffer bytes.Buffer
   keys := make([]string, 0, len(list))
   for k := range list {
     keys = append(keys, k)
   }

   sort.Strings(keys)

   for _, key := range keys {
     value := list[key]
     buffer.WriteString(encodeUrl(key))
     buffer.WriteString("=")
     buffer.WriteString(encodeUrl(value))
     buffer.WriteString("&")
   }
   len := buffer.Len() - 1
   buffer.Truncate(len)
   return buffer.String(), nil
}
