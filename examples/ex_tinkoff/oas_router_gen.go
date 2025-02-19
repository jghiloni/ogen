// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/jghiloni/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'm': // Prefix: "market/"
				origElem := elem
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					origElem := elem
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketBondsGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'c': // Prefix: "c"
					origElem := elem
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "andles"
						origElem := elem
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketCandlesGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					case 'u': // Prefix: "urrencies"
						origElem := elem
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketCurrenciesGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'e': // Prefix: "etfs"
					origElem := elem
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketEtfsGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'o': // Prefix: "orderbook"
					origElem := elem
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketOrderbookGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 's': // Prefix: "s"
					origElem := elem
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						origElem := elem
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							origElem := elem
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleMarketSearchByFigiGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

							elem = origElem
						case 't': // Prefix: "ticker"
							origElem := elem
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleMarketSearchByTickerGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					case 't': // Prefix: "tocks"
						origElem := elem
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketStocksGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'o': // Prefix: "o"
				origElem := elem
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					origElem := elem
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOperationsGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'r': // Prefix: "rders"
					origElem := elem
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleOrdersGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "cancel"
							origElem := elem
							if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleOrdersCancelPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						case 'l': // Prefix: "limit-order"
							origElem := elem
							if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleOrdersLimitOrderPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						case 'm': // Prefix: "market-order"
							origElem := elem
							if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleOrdersMarketOrderPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "portfolio"
				origElem := elem
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handlePortfolioGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					origElem := elem
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handlePortfolioCurrenciesGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "sandbox/"
				origElem := elem
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "c"
					origElem := elem
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "lear"
						origElem := elem
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxClearPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'u': // Prefix: "urrencies/balance"
						origElem := elem
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxCurrenciesBalancePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "positions/balance"
					origElem := elem
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleSandboxPositionsBalancePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'r': // Prefix: "re"
					origElem := elem
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						origElem := elem
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxRegisterPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'm': // Prefix: "move"
						origElem := elem
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxRemovePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'u': // Prefix: "user/accounts"
				origElem := elem
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleUserAccountsGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'm': // Prefix: "market/"
				origElem := elem
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					origElem := elem
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "MarketBondsGet"
							r.summary = "Получение списка облигаций"
							r.operationID = ""
							r.pathPattern = "/market/bonds"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'c': // Prefix: "c"
					origElem := elem
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "andles"
						origElem := elem
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "MarketCandlesGet"
								r.summary = "Получение исторических свечей по FIGI"
								r.operationID = ""
								r.pathPattern = "/market/candles"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'u': // Prefix: "urrencies"
						origElem := elem
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "MarketCurrenciesGet"
								r.summary = "Получение списка валютных пар"
								r.operationID = ""
								r.pathPattern = "/market/currencies"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'e': // Prefix: "etfs"
					origElem := elem
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "MarketEtfsGet"
							r.summary = "Получение списка ETF"
							r.operationID = ""
							r.pathPattern = "/market/etfs"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'o': // Prefix: "orderbook"
					origElem := elem
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "MarketOrderbookGet"
							r.summary = "Получение стакана по FIGI"
							r.operationID = ""
							r.pathPattern = "/market/orderbook"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 's': // Prefix: "s"
					origElem := elem
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						origElem := elem
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							origElem := elem
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = "MarketSearchByFigiGet"
									r.summary = "Получение инструмента по FIGI"
									r.operationID = ""
									r.pathPattern = "/market/search/by-figi"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 't': // Prefix: "ticker"
							origElem := elem
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = "MarketSearchByTickerGet"
									r.summary = "Получение инструмента по тикеру"
									r.operationID = ""
									r.pathPattern = "/market/search/by-ticker"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					case 't': // Prefix: "tocks"
						origElem := elem
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "MarketStocksGet"
								r.summary = "Получение списка акций"
								r.operationID = ""
								r.pathPattern = "/market/stocks"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'o': // Prefix: "o"
				origElem := elem
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					origElem := elem
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "OperationsGet"
							r.summary = "Получение списка операций"
							r.operationID = ""
							r.pathPattern = "/operations"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'r': // Prefix: "rders"
					origElem := elem
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "OrdersGet"
							r.summary = "Получение списка активных заявок"
							r.operationID = ""
							r.pathPattern = "/orders"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "cancel"
							origElem := elem
							if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "OrdersCancelPost"
									r.summary = "Отмена заявки"
									r.operationID = ""
									r.pathPattern = "/orders/cancel"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'l': // Prefix: "limit-order"
							origElem := elem
							if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "OrdersLimitOrderPost"
									r.summary = "Создание лимитной заявки"
									r.operationID = ""
									r.pathPattern = "/orders/limit-order"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'm': // Prefix: "market-order"
							origElem := elem
							if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "OrdersMarketOrderPost"
									r.summary = "Создание рыночной заявки"
									r.operationID = ""
									r.pathPattern = "/orders/market-order"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "portfolio"
				origElem := elem
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "PortfolioGet"
						r.summary = "Получение портфеля клиента"
						r.operationID = ""
						r.pathPattern = "/portfolio"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					origElem := elem
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "PortfolioCurrenciesGet"
							r.summary = "Получение валютных активов клиента"
							r.operationID = ""
							r.pathPattern = "/portfolio/currencies"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "sandbox/"
				origElem := elem
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "c"
					origElem := elem
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "lear"
						origElem := elem
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "SandboxClearPost"
								r.summary = "Удаление всех позиций"
								r.operationID = ""
								r.pathPattern = "/sandbox/clear"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'u': // Prefix: "urrencies/balance"
						origElem := elem
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "SandboxCurrenciesBalancePost"
								r.summary = "Выставление баланса по валютным позициям"
								r.operationID = ""
								r.pathPattern = "/sandbox/currencies/balance"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "positions/balance"
					origElem := elem
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "SandboxPositionsBalancePost"
							r.summary = "Выставление баланса по инструментным позициям"
							r.operationID = ""
							r.pathPattern = "/sandbox/positions/balance"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'r': // Prefix: "re"
					origElem := elem
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						origElem := elem
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "SandboxRegisterPost"
								r.summary = "Регистрация клиента в sandbox"
								r.operationID = ""
								r.pathPattern = "/sandbox/register"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'm': // Prefix: "move"
						origElem := elem
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "SandboxRemovePost"
								r.summary = "Удаление счета"
								r.operationID = ""
								r.pathPattern = "/sandbox/remove"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'u': // Prefix: "user/accounts"
				origElem := elem
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = "UserAccountsGet"
						r.summary = "Получение брокерских счетов клиента"
						r.operationID = ""
						r.pathPattern = "/user/accounts"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
