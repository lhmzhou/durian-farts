# durian-farts

`durian-farts` is a reusable Go package for transforming keys in a JSON object (input as a string) using a supplied string transformation function. Out-of-the-box, `durian-farts` implements the transformation functions carried out in [strcase](https://github.com/iancoleman/strcase). 

### Sample Input

```go
in := `{
  "firstName": "Higgs",
  "lastName": "Boson",
  "address": {
    "streetAddress": "Esplanade des Particules 1",
    "city": "Meyrin",
    "canton": "Geneva",
    "countryCode": "Switzerland"
  },
  "lineOfBusiness": [
    { "reasonCode": "SR-042", "description": "Standard Model of contemporary particle physics" }
  ],
  "futureUse": [
    { "useCode": "QW-011", "description": " Validation of LHC, etc." }
  ]
}`
```

## Examples of Outputs

#### Snake Case

```json
{
  "address": {
    "city": "Meyrin",
    "country_code": "Switzerland",
    "canton": "Geneva",
    "street_address": "Esplanade des Particules 1"
  },
  "first_name": "Higgs",
  "future_use": [
    {
      "description": " Validation of LHC, etc.",
      "use_code": "QW-011"
    }
  ],
  "last_name": "Boson",
  "line_of_business": [
    {
      "description": "Standard Model of contemporary particle physics",
      "reason_code": "SR-042"
    }
  ]
}
```

#### Screaming Snake Case

```json
{
  "ADDRESS": {
    "CITY": "Meyrin",
    "COUNTRY_CODE": "Switzerland",
    "canton": "Geneva",
    "STREET_ADDRESS": "Esplanade des Particules 1"
  },
  "FIRST_NAME": "Higgs",
  "FUTURE_USE": [
    {
      "DESCRIPTION": " Validation of LHC, etc.",
      "USE_CODE": "QW-011"
    }
  ],
  "LAST_NAME": "Boson",
  "LINE_OF_BUSINESS": [
    {
      "DESCRIPTION": "Standard Model of contemporary particle physics",
      "REASON_CODE": "SR-042"
    }
  ]
}
```

#### Kebab Case

```json
{
  "address": {
    "city": "Meyrin",
    "country-code": "Switzerland",
    "canton": "Geneva",
    "street-address": "Esplanade des Particules 1"
  },
  "first-name": "Higgs",
  "future-use": [
    {
      "description": " Validation of LHC, etc.",
      "use-code": "QW-011"
    }
  ],
  "last-name": "Boson",
  "line-of-business": [
    {
      "description": "Standard Model of contemporary particle physics",
      "reason-code": "SR-042"
    }
  ]
}
```

#### Screaming Kebab Case

```json
{
  "ADDRESS": {
    "CITY": "Meyrin",
    "COUNTRY-CODE": "Switzerland",
    "canton": "Geneva",
    "STREET-ADDRESS": "Esplanade des Particules 1"
  },
  "FIRST-NAME": "Higgs",
  "FUTURE-USE": [
    {
      "DESCRIPTION": " Validation of LHC, etc.",
      "USE-CODE": "QW-011"
    }
  ],
  "LAST-NAME": "Boson",
  "LINE-OF-BUSINESS": [
    {
      "DESCRIPTION": "Standard Model of contemporary particle physics",
      "REASON-CODE": "SR-042"
    }
  ]
}
```

#### Camel Case

```json
{
  "Address": {
    "City": "Meyrin",
    "CountryCode": "Switzerland",
    "canton": "Geneva",
    "StreetAddress": "Esplanade des Particules 1"
  },
  "FirstName": "Higgs",
  "FutureUse": [
    {
      "Description": " Validation of LHC, etc.",
      "UseCode": "QW-011"
    }
  ],
  "LastName": "Boson",
  "lineOfBusiness": [
    {
      "Description": "Standard Model of contemporary particle physics",
      "ReasonCode": "SR-042"
    }
  ]
}
```

#### Lower Camel Case

```json
{
  "address": {
    "city": "Meyrin",
    "countryCode": "Switzerland",
    "canton": "Geneva",
    "streetAddress": "Esplanade des Particules 1"
  },
  "firstName": "Higgs",
  "futureUse": [
    {
      "description": " Validation of LHC, etc.",
      "useCode": "QW-011"
    }
  ],
  "lastName": "Boson",
  "lineOfBusiness": [
    {
      "description": "Standard Model of contemporary particle physics",
      "reasonCode": "SR-042"
    }
  ]
}
```

#### Custom Case

Using the following custom transformation function:

```go
func(s string) string {
  var o string = s
  o = strings.Replace(o, "a", "4", -1)
  o = strings.Replace(o, "i", "1", -1)
  return o
})
```

returns:

```json
{
  "4ddress": {
    "c1ty": "Meyrin",
    "countryCode": "Switzerland",
    "prov1nce": "Geneva",
    "streetAddress": "Esplanade des Particules 1"
  },
  "f1rstN4me": "Higgs",
  "futureUse": [
    {
      "descr1pt1on": " Validation of LHC, etc.",
      "useCode": "QW-011"
    }
  ],
  "l4stN4me": "Boson",
  "l1neOfBus1ness": [
    {
      "descr1pt1on": "Standard Model of contemporary particle physics",
      "re4sonCode": "SR-042"
    }
  ]
}
```
