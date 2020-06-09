package jwt

import (
	"testing"
)

func TestNewGenerator(t *testing.T) {
	g := NewGenerator("aaaaaaaa-bbbb-cccc-dddd-0123456789ab", []byte(getPrivateKey()))

	if g.ApplicationID != "aaaaaaaa-bbbb-cccc-dddd-0123456789ab" {
		t.Errorf("Application ID not added to token generator")
	}
}

func TestOnePath(t *testing.T) {
	g := NewGenerator("aaaaaaaa-bbbb-cccc-dddd-0123456789ab", []byte(getPrivateKey()))
	g.AddPath(Path{Path: "/*/users/**"})

	token, _ := g.GenerateToken()
	// start with checking we got vaguely a token
	if len(token) < 400 {
		t.Errorf("Token may not have been generated")
	}

}

func getPrivateKey() string {
	return `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCxeDzyjtX+wNSf
SfRhOnvbfj3wa9NFzpUGF1OsLytdij0B9+RRGnHezfJE39O5A2W75DuBiXZEhjw2
YlPIsrpxqdKDVoOtYqIkmojbjs47yYIkpdAWDybwA4oE47TiCZEentrn6nHZgBuJ
ueSb1C7pwcSElzd+adEWiE9BnzG7urmOF6actUZTlhBVjFWbMzU9dtXaXKhfLEKT
/O6khOPEgcBa5ogoPsc5X/q9X0BhiAR/vb4dONeka/P4PtYiYIRye4sqatpyiKmT
5fteSWId/Kw+MreeGq8y04aj/n+aCZsDuWCzuLSbsmKk9kpXOUAiJ025XhNpuqlk
5TDxV1idAgMBAAECggEAUzgCQG2ZTG4I5i7sMSGXDdx7V/9/4TaXa/VJT68Iw0K9
C+y9vAhOCEADkKdypUnCDWLfQoV+l3Bo0Mm35x9kTUNoixzo/0eGp+ptLaOf8qox
5FR6qLVvOVI3iMZsPihlS/oKxYCK5YZso18vo6DSZNvvdotgQt+E4++EVs27KP+E
Y3A6oXxAysG2O1Ruip840S1XpcdlcjNXNoI0I5u7yiVdnXle+dvQaBvR85+S7dJl
XOo9mJh40248BwbHO7B0fmKgT724pLZTlJwyElH7QE9o8px27/Qb2nUMnyE8Txqv
HwaKWXgxEnM+M29n+brA0JHhqG7VtLRl7wkaIHVV0wKBgQDmAA7gd1Iz4EfOI7nm
b3sbpv9AWttLRqJ4FGbtZYrycMCPzyPkJW75NJAxKFgUh3R+uE3GT/1qPIyKiGlk
21vV9X2nsXajG3aUuDaJMZpLj2cwGBg5FND1ebM1BuSVxcf+FKx5BZ2qhtKywm/o
qihZK8ou7lph+mKjjpME7IDHkwKBgQDFiAGianyymlY+QQtLECGmALFUmsyFac0f
9qaBe58RuF+/uh0pTyW8ryvHCai9UsHuRUsnCh8MBr79szOq5MQYjdWeYOzZW/PP
qvwOti6ZHj/oEQyy9HBs043+ZtOvn4Jqgl5xMwnyB/FSI6mCxEWoNlrV/h7Pm00w
LiNHohAdDwKBgQCjBLmGqawO7tGWL0ZNKQj10YiroXo7QYZYXAWUD9vK+NXTWDsV
Mt8ULQhifzjm3Bda1eDyRbHVQbNPYV5qSPwvi3+TgzoWY5nJ0UN+PDUjhzHZKcrJ
cpKk2qyFUixkZ7nXwel5IdzdiBAA+cV+AFT21w3H89MGDQUq7hwQalzglwKBgBL4
at/ERlGIzPuRl5oP5ItiyaUMcNPnQ1HsiDUrQC2dfWSWZTKQbVlfoV6uKMx15DYT
5ZHMQQsQossOqMmiyspo5LkfKd/+Gr4495gaGwONiimpaYTOQPXSo3JpQa8+LHI6
LwPEGRJrfNucnSz32JC9F6AxlZfunE0iQTUh9VY1AoGAdk83WFbgzO95MYdW/9oe
AW78Q3WCS/+tDi6aud+S5D5MiWGvfBBcMD12616nF0mZI8340FGar2ep9hnOaZ0t
DdpUM51nYx7B6mK9sDrv+7H2xLIWI3ttr9waABDp1mBduRThleTOPaInxIzAjFgb
d7d9YI+xhf0OOEYNLiSTLdI=
-----END PRIVATE KEY-----`
}
