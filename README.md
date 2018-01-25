```
___________
\__    ___/  ____  ______  _______   ____  ______    ____
  |    |    /  _ \ \____ \ \_  __ \ /  _ \ \____ \ _/ __ \
  |    |   (  <_> )|  |_> > |  | \/(  <_> )|  |_> >\  ___/
  |____|    \____/ |   __/  |__|    \____/ |   __/  \___  >
                   |__|                    |__|         \/
```
# Toprope
[![Build Status](https://travis-ci.org/deadcheat/toprope.svg?branch=master)](https://travis-ci.org/deadcheat/toprope) [![Coverage Status](https://coveralls.io/repos/github/deadcheat/toprope/badge.svg?branch=master)](https://coveralls.io/github/deadcheat/toprope?branch=master)

## What's this

toprope is aiming to flexible wrapper of net/http/test

httptest.NewServer assigns random port of localhost.
Ofcourse, it's good, yes.

but i want to specify a port or sometimes hostname too.

that's why i created this lib. thank you.

## How To Use

```
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	
	ts, err := toprope.NewHttptestTCPServerFromURL("http://localhost:8080", h)
	if err != nil {
		t.Error("failed to start local-http-server")
		t.Fail()
	}
	ts.Start()
	defer ts.Close()
```

## Notice
- toprope can't create TLS server now


thankyou.
