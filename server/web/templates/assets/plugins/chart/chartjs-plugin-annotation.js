/*
 * chartjs-plugin-annotation.js
 * http://chartjs.org/
 * Version: 0.5.7
 *
 * Copyright 2016 Evert Timberg
 * Released under the MIT license
 * https://github.com/chartjs/Chart.Annotation.js/blob/master/LICENSE.md
 */
(function e(g, b, d) {
    function f(k, m) {
        if (!b[k]) {
            if (!g[k]) {
                var h = typeof require == "function" && require;
                if (!m && h) {
                    return h(k, !0)
                }
                if (a) {
                    return a(k, !0)
                }
                var i = new Error("Cannot find module '" + k + "'");
                throw i.code = "MODULE_NOT_FOUND", i
            }
            var j = b[k] = {
                exports: {}
            };
            g[k][0].call(j.exports, function(l) {
                var o = g[k][1][l];
                return f(o ? o : l)
            }, j, j.exports, e, g, b, d)
        }
        return b[k].exports
    }
    var a = typeof require == "function" && require;
    for (var c = 0; c < d.length; c++) {
        f(d[c])
    }
    return f
})({
    1: [function(c, b, a) {}, {}],
    2: [function(c, b, a) {
        b.exports = function(f) {
            var g = f.helpers;
            var j = c("./helpers.js")(f);
            var i = c("./events.js")(f);
            var d = f.Annotation.types;

            function k(l) {
                j.decorate(l, "afterDataLimits", function(m, n) {
                    if (m) {
                        m(n)
                    }
                    j.adjustScaleRange(n)
                })
            }

            function h(l) {
                return function(m, o) {
                    var n = m.annotation.options.drawTime;
                    j.elements(m).filter(function(p) {
                        return l === (p.options.drawTime || n)
                    }).forEach(function(p) {
                        p.transition(o).draw()
                    })
                }
            }
            return {
                beforeInit: function(l) {
                    var m = l.options;
                    var n = l.annotation = {
                        elements: {},
                        options: j.initConfig(m.annotation || {}),
                        onDestroy: [],
                        firstRun: true,
                        supported: false
                    };
                    l.ensureScalesHaveIDs();
                    if (m.scales) {
                        n.supported = true;
                        g.each(m.scales.xAxes, k);
                        g.each(m.scales.yAxes, k)
                    }
                },
                beforeUpdate: function(l) {
                    var n = l.annotation;
                    if (!n.supported) {
                        return
                    }
                    if (!n.firstRun) {
                        n.options = j.initConfig(l.options.annotation || {})
                    } else {
                        n.firstRun = false
                    }
                    var m = [];
                    n.options.annotations.forEach(function(o) {
                        var r = o.id || j.objectId();
                        if (!n.elements[r] && d[o.type]) {
                            var p = d[o.type];
                            var q = new p({
                                id: r,
                                options: o,
                                chartInstance: l,
                            });
                            q.initialize();
                            n.elements[r] = q;
                            o.id = r;
                            m.push(r)
                        } else {
                            if (n.elements[r]) {
                                m.push(r)
                            }
                        }
                    });
                    Object.keys(n.elements).forEach(function(o) {
                        if (m.indexOf(o) === -1) {
                            n.elements[o].destroy();
                            delete n.elements[o]
                        }
                    })
                },
                afterScaleUpdate: function(l) {
                    j.elements(l).forEach(function(m) {
                        m.configure()
                    })
                },
                beforeDatasetsDraw: h("beforeDatasetsDraw"),
                afterDatasetsDraw: h("afterDatasetsDraw"),
                afterDraw: h("afterDraw"),
                afterInit: function(m) {
                    var o = m.annotation.options.events;
                    if (g.isArray(o) && o.length > 0) {
                        var l = m.chart.canvas;
                        var n = i.dispatcher.bind(m);
                        i.collapseHoverEvents(o).forEach(function(p) {
                            g.addEvent(l, p, n);
                            m.annotation.onDestroy.push(function() {
                                g.removeEvent(l, p, n)
                            })
                        })
                    }
                },
                destroy: function(l) {
                    var m = l.annotation.onDestroy;
                    while (m.length > 0) {
                        m.pop()()
                    }
                }
            }
        }
    }, {
        "./events.js": 4,
        "./helpers.js": 5
    }],
    3: [function(c, b, a) {
        b.exports = function(f) {
            var g = f.helpers;
            var d = f.Element.extend({
                initialize: function() {
                    this.hidden = false;
                    this.hovering = false;
                    this._model = g.clone(this._model) || {};
                    this.setDataLimits()
                },
                destroy: function() {},
                setDataLimits: function() {},
                configure: function() {},
                inRange: function() {},
                getCenterPoint: function() {},
                getWidth: function() {},
                getHeight: function() {},
                getArea: function() {},
                draw: function() {}
            });
            return d
        }
    }, {}],
    4: [function(c, b, a) {
        b.exports = function(d) {
            var f = d.helpers;
            var i = c("./helpers.js")(d);

            function g(j) {
                var l = false;
                var k = j.filter(function(m) {
                    switch (m) {
                        case "mouseenter":
                        case "mouseover":
                        case "mouseout":
                        case "mouseleave":
                            l = true;
                            return false;
                        default:
                            return true
                    }
                });
                if (l && k.indexOf("mousemove") === -1) {
                    k.push("mousemove")
                }
                return k
            }

            function h(k) {
                var q = this.annotation;
                var m = i.elements(this);
                var s = f.getRelativePosition(k, this.chart);
                var l = i.getNearestItems(m, s);
                var p = g(q.options.events);
                var j = q.options.dblClickSpeed;
                var o = [];
                var n = i.getEventHandlerName(k.type);
                var r = (l || {}).options;
                if (k.type === "mousemove") {
                    if (l && !l.hovering) {
                        ["mouseenter", "mouseover"].forEach(function(u) {
                            var t = i.getEventHandlerName(u);
                            var v = i.createMouseEvent(u, k);
                            l.hovering = true;
                            if (typeof r[t] === "function") {
                                o.push([r[t], v, l])
                            }
                        })
                    } else {
                        if (!l) {
                            m.forEach(function(t) {
                                if (t.hovering) {
                                    t.hovering = false;
                                    var u = t.options;
                                    ["mouseout", "mouseleave"].forEach(function(w) {
                                        var v = i.getEventHandlerName(w);
                                        var x = i.createMouseEvent(w, k);
                                        if (typeof u[v] === "function") {
                                            o.push([u[v], x, t])
                                        }
                                    })
                                }
                            })
                        }
                    }
                }
                if (l && p.indexOf("dblclick") > -1 && typeof r.onDblclick === "function") {
                    if (k.type === "click" && typeof r.onClick === "function") {
                        clearTimeout(l.clickTimeout);
                        l.clickTimeout = setTimeout(function() {
                            delete l.clickTimeout;
                            r.onClick.call(l, k)
                        }, j);
                        k.stopImmediatePropagation();
                        k.preventDefault();
                        return
                    } else {
                        if (k.type === "dblclick" && l.clickTimeout) {
                            clearTimeout(l.clickTimeout);
                            delete l.clickTimeout
                        }
                    }
                }
                if (l && typeof r[n] === "function" && o.length === 0) {
                    o.push([r[n], k, l])
                }
                if (o.length > 0) {
                    k.stopImmediatePropagation();
                    k.preventDefault();
                    o.forEach(function(t) {
                        t[0].call(t[2], t[1])
                    })
                }
            }
            return {
                dispatcher: h,
                collapseHoverEvents: g
            }
        }
    }, {
        "./helpers.js": 5
    }],
    5: [function(l, i, f) {
        function j() {}

        function d(m) {
            var n = m.annotation.elements;
            return Object.keys(n).map(function(o) {
                return n[o]
            })
        }

        function k() {
            return Math.random().toString(36).substr(2, 6)
        }

        function h(m) {
            if (m === null || typeof m === "undefined") {
                return false
            } else {
                if (typeof m === "number") {
                    return isFinite(m)
                } else {
                    return !!m
                }
            }
        }

        function c(n, p, m) {
            var o = "$";
            if (!n[o + p]) {
                if (n[p]) {
                    n[o + p] = n[p].bind(n);
                    n[p] = function() {
                        var q = [n[o + p]].concat(Array.prototype.slice.call(arguments));
                        return m.apply(n, q)
                    }
                } else {
                    n[p] = function() {
                        var q = [undefined].concat(Array.prototype.slice.call(arguments));
                        return m.apply(n, q)
                    }
                }
            }
        }

        function a(m, n) {
            m.forEach(function(o) {
                (n ? o[n] : o)()
            })
        }

        function g(m) {
            return "on" + m[0].toUpperCase() + m.substring(1)
        }

        function b(s, r) {
            try {
                return new MouseEvent(s, r)
            } catch (o) {
                try {
                    var q = document.createEvent("MouseEvent");
                    q.initMouseEvent(s, r.canBubble, r.cancelable, r.view, r.detail, r.screenX, r.screenY, r.clientX, r.clientY, r.ctrlKey, r.altKey, r.shiftKey, r.metaKey, r.button, r.relatedTarget);
                    return q
                } catch (p) {
                    var n = document.createEvent("Event");
                    n.initEvent(s, r.canBubble, r.cancelable);
                    return n
                }
            }
        }
        i.exports = function(n) {
            var o = n.helpers;

            function r(s) {
                s = o.configMerge(n.Annotation.defaults, s);
                if (o.isArray(s.annotations)) {
                    s.annotations.forEach(function(t) {
                        t.label = o.configMerge(n.Annotation.labelDefaults, t.label)
                    })
                }
                return s
            }

            function q(w, s, y, x) {
                var v = s.filter(function(z) {
                    return !!z._model.ranges[w]
                }).map(function(z) {
                    return z._model.ranges[w]
                });
                var u = v.map(function(z) {
                    return Number(z.min)
                }).reduce(function(z, A) {
                    return isFinite(A) && !isNaN(A) && A < z ? A : z
                }, y);
                var t = v.map(function(z) {
                    return Number(z.max)
                }).reduce(function(z, A) {
                    return isFinite(A) && !isNaN(A) && A > z ? A : z
                }, x);
                return {
                    min: u,
                    max: t
                }
            }

            function m(t) {
                var s = q(t.id, d(t.chart), t.min, t.max);
                if (typeof t.options.ticks.min === "undefined" && typeof t.options.ticks.suggestedMin === "undefined") {
                    t.min = s.min
                }
                if (typeof t.options.ticks.max === "undefined" && typeof t.options.ticks.suggestedMax === "undefined") {
                    t.max = s.max
                }
                if (t.handleTickRangeOptions) {
                    t.handleTickRangeOptions()
                }
            }

            function p(s, u) {
                var t = Number.POSITIVE_INFINITY;
                return s.filter(function(v) {
                    return v.inRange(u.x, u.y)
                }).reduce(function(y, x) {
                    var v = x.getCenterPoint();
                    var w = o.distanceBetweenPoints(u, v);
                    if (w < t) {
                        y = [x];
                        t = w
                    } else {
                        if (w === t) {
                            y.push(x)
                        }
                    }
                    return y
                }, []).sort(function(v, w) {
                    var x = v.getArea(),
                        y = w.getArea();
                    return (x > y || x < y) ? x - y : v._index - w._index
                }).slice(0, 1)[0]
            }
            return {
                initConfig: r,
                elements: d,
                callEach: a,
                noop: j,
                objectId: k,
                isValid: h,
                decorate: c,
                adjustScaleRange: m,
                getNearestItems: p,
                getEventHandlerName: g,
                createMouseEvent: b
            }
        }
    }, {}],
    6: [function(f, d, c) {
        var b = f("chart.js");
        b = typeof(b) === "function" ? b : window.Chart;
        b.Annotation = b.Annotation || {};
        b.Annotation.drawTimeOptions = {
            afterDraw: "afterDraw",
            afterDatasetsDraw: "afterDatasetsDraw",
            beforeDatasetsDraw: "beforeDatasetsDraw"
        };
        b.Annotation.defaults = {
            drawTime: "afterDatasetsDraw",
            dblClickSpeed: 350,
            events: [],
            annotations: []
        };
        b.Annotation.labelDefaults = {
            backgroundColor: "rgba(0,0,0,0.8)",
            fontFamily: b.defaults.global.defaultFontFamily,
            fontSize: b.defaults.global.defaultFontSize,
            fontStyle: "bold",
            fontColor: "#fff",
            xPadding: 6,
            yPadding: 6,
            cornerRadius: 6,
            position: "center",
            xAdjust: 0,
            yAdjust: 0,
            enabled: false,
            content: null
        };
        b.Annotation.Element = f("./element.js")(b);
        b.Annotation.types = {
            line: f("./types/line.js")(b),
            box: f("./types/box.js")(b)
        };
        var a = f("./annotation.js")(b);
        d.exports = a;
        b.pluginService.register(a)
    }, {
        "./annotation.js": 2,
        "./element.js": 3,
        "./types/box.js": 7,
        "./types/line.js": 8,
        "chart.js": 1
    }],
    7: [function(c, b, a) {
        b.exports = function(f) {
            var g = c("../helpers.js")(f);
            var d = f.Annotation.Element.extend({
                setDataLimits: function() {
                    var l = this._model;
                    var m = this.options;
                    var i = this.chartInstance;
                    var n = i.scales[m.xScaleID];
                    var o = i.scales[m.yScaleID];
                    var h = i.chartArea;
                    l.ranges = {};
                    if (!h) {
                        return
                    }
                    var k = 0;
                    var j = 0;
                    if (n) {
                        k = g.isValid(m.xMin) ? m.xMin : n.getPixelForValue(h.left);
                        j = g.isValid(m.xMax) ? m.xMax : n.getPixelForValue(h.right);
                        l.ranges[m.xScaleID] = {
                            min: Math.min(k, j),
                            max: Math.max(k, j)
                        }
                    }
                    if (o) {
                        k = g.isValid(m.yMin) ? m.yMin : o.getPixelForValue(h.bottom);
                        j = g.isValid(m.yMax) ? m.yMax : o.getPixelForValue(h.top);
                        l.ranges[m.yScaleID] = {
                            min: Math.min(k, j),
                            max: Math.max(k, j)
                        }
                    }
                },
                configure: function() {
                    var n = this._model;
                    var o = this.options;
                    var j = this.chartInstance;
                    var r = j.scales[o.xScaleID];
                    var s = j.scales[o.yScaleID];
                    var i = j.chartArea;
                    n.clip = {
                        x1: i.left,
                        x2: i.right,
                        y1: i.top,
                        y2: i.bottom
                    };
                    var k = i.left,
                        q = i.top,
                        p = i.right,
                        h = i.bottom;
                    var m, l;
                    if (r) {
                        m = g.isValid(o.xMin) ? r.getPixelForValue(o.xMin) : i.left;
                        l = g.isValid(o.xMax) ? r.getPixelForValue(o.xMax) : i.right;
                        k = Math.min(m, l);
                        p = Math.max(m, l)
                    }
                    if (s) {
                        m = g.isValid(o.yMin) ? s.getPixelForValue(o.yMin) : i.bottom;
                        l = g.isValid(o.yMax) ? s.getPixelForValue(o.yMax) : i.top;
                        q = Math.min(m, l);
                        h = Math.max(m, l)
                    }
                    n.left = k;
                    n.top = q;
                    n.right = p;
                    n.bottom = h;
                    n.borderColor = o.borderColor;
                    n.borderWidth = o.borderWidth;
                    n.backgroundColor = o.backgroundColor
                },
                inRange: function(i, j) {
                    var h = this._model;
                    return h && i >= h.left && i <= h.right && j >= h.top && j <= h.bottom
                },
                getCenterPoint: function() {
                    var h = this._model;
                    return {
                        x: (h.right + h.left) / 2,
                        y: (h.bottom + h.top) / 2
                    }
                },
                getWidth: function() {
                    var h = this._model;
                    return Math.abs(h.right - h.left)
                },
                getHeight: function() {
                    var h = this._model;
                    return Math.abs(h.bottom - h.top)
                },
                getArea: function() {
                    return this.getWidth() * this.getHeight()
                },
                draw: function() {
                    var j = this._view;
                    var h = this.chartInstance.chart.ctx;
                    h.save();
                    h.beginPath();
                    h.rect(j.clip.x1, j.clip.y1, j.clip.x2 - j.clip.x1, j.clip.y2 - j.clip.y1);
                    h.clip();
                    h.lineWidth = j.borderWidth;
                    h.strokeStyle = j.borderColor;
                    h.fillStyle = j.backgroundColor;
                    var k = j.right - j.left,
                        i = j.bottom - j.top;
                    h.fillRect(j.left, j.top, k, i);
                    h.strokeRect(j.left, j.top, k, i);
                    h.restore()
                }
            });
            return d
        }
    }, {
        "../helpers.js": 5
    }],
    8: [function(c, b, a) {
        b.exports = function(f) {
            var g = f.helpers;
            var h = c("../helpers.js")(f);
            var i = "horizontal";
            var l = "vertical";
            var j = f.Annotation.Element.extend({
                setDataLimits: function() {
                    var m = this._model;
                    var n = this.options;
                    m.ranges = {};
                    m.ranges[n.scaleID] = {
                        min: n.value,
                        max: n.endValue || n.value
                    }
                },
                configure: function() {
                    var r = this._model;
                    var s = this.options;
                    var n = this.chartInstance;
                    var o = n.chart.ctx;
                    var u = n.scales[s.scaleID];
                    var t, p;
                    if (u) {
                        t = h.isValid(s.value) ? u.getPixelForValue(s.value) : NaN;
                        p = h.isValid(s.endValue) ? u.getPixelForValue(s.endValue) : t
                    }
                    if (isNaN(t)) {
                        return
                    }
                    var m = n.chartArea;
                    r.clip = {
                        x1: m.left,
                        x2: m.right,
                        y1: m.top,
                        y2: m.bottom
                    };
                    if (this.options.mode == i) {
                        r.x1 = m.left;
                        r.x2 = m.right;
                        r.y1 = t;
                        r.y2 = p
                    } else {
                        r.y1 = m.top;
                        r.y2 = m.bottom;
                        r.x1 = t;
                        r.x2 = p
                    }
                    r.line = new k(r);
                    r.mode = s.mode;
                    r.labelBackgroundColor = s.label.backgroundColor;
                    r.labelFontFamily = s.label.fontFamily;
                    r.labelFontSize = s.label.fontSize;
                    r.labelFontStyle = s.label.fontStyle;
                    r.labelFontColor = s.label.fontColor;
                    r.labelXPadding = s.label.xPadding;
                    r.labelYPadding = s.label.yPadding;
                    r.labelCornerRadius = s.label.cornerRadius;
                    r.labelPosition = s.label.position;
                    r.labelXAdjust = s.label.xAdjust;
                    r.labelYAdjust = s.label.yAdjust;
                    r.labelEnabled = s.label.enabled;
                    r.labelContent = s.label.content;
                    o.font = g.fontString(r.labelFontSize, r.labelFontStyle, r.labelFontFamily);
                    var w = o.measureText(r.labelContent).width;
                    var v = o.measureText("M").width;
                    var q = d(r, w, v, r.labelXPadding, r.labelYPadding);
                    r.labelX = q.x - r.labelXPadding;
                    r.labelY = q.y - r.labelYPadding;
                    r.labelWidth = w + (2 * r.labelXPadding);
                    r.labelHeight = v + (2 * r.labelYPadding);
                    r.borderColor = s.borderColor;
                    r.borderWidth = s.borderWidth;
                    r.borderDash = s.borderDash || [];
                    r.borderDashOffset = s.borderDashOffset || 0
                },
                inRange: function(n, o) {
                    var m = this._model;
                    return (m.line && m.line.intersects(n, o, this.getHeight())) || (m.labelEnabled && m.labelContent && n >= m.labelX && n <= m.labelX + m.labelWidth && o >= m.labelY && o <= m.labelY + m.labelHeight)
                },
                getCenterPoint: function() {
                    return {
                        x: (this._model.x2 + this._model.x1) / 2,
                        y: (this._model.y2 + this._model.y1) / 2
                    }
                },
                getWidth: function() {
                    return Math.abs(this._model.right - this._model.left)
                },
                getHeight: function() {
                    return this._model.borderWidth || 1
                },
                getArea: function() {
                    return Math.sqrt(Math.pow(this.getWidth(), 2) + Math.pow(this.getHeight(), 2))
                },
                draw: function() {
                    var n = this._view;
                    var m = this.chartInstance.chart.ctx;
                    if (!n.clip) {
                        return
                    }
                    m.save();
                    m.beginPath();
                    m.rect(n.clip.x1, n.clip.y1, n.clip.x2 - n.clip.x1, n.clip.y2 - n.clip.y1);
                    m.clip();
                    m.lineWidth = n.borderWidth;
                    m.strokeStyle = n.borderColor;
                    if (m.setLineDash) {
                        m.setLineDash(n.borderDash)
                    }
                    m.lineDashOffset = n.borderDashOffset;
                    m.beginPath();
                    m.moveTo(n.x1, n.y1);
                    m.lineTo(n.x2, n.y2);
                    m.stroke();
                    if (n.labelEnabled && n.labelContent) {
                        m.beginPath();
                        m.rect(n.clip.x1, n.clip.y1, n.clip.x2 - n.clip.x1, n.clip.y2 - n.clip.y1);
                        m.clip();
                        m.fillStyle = n.labelBackgroundColor;
                        g.drawRoundedRectangle(m, n.labelX, n.labelY, n.labelWidth, n.labelHeight, n.labelCornerRadius);
                        m.fill();
                        m.font = g.fontString(n.labelFontSize, n.labelFontStyle, n.labelFontFamily);
                        m.fillStyle = n.labelFontColor;
                        m.textAlign = "center";
                        m.textBaseline = "middle";
                        m.fillText(n.labelContent, n.labelX + (n.labelWidth / 2), n.labelY + (n.labelHeight / 2))
                    }
                    m.restore()
                }
            });

            function k(p) {
                var o = (p.x2 - p.x1) / (p.y2 - p.y1);
                var n = p.x1 || 0;
                this.m = o;
                this.b = n;
                this.getX = function(m) {
                    return o * (m - p.y1) + n
                };
                this.getY = function(m) {
                    return ((m - n) / o) + p.y1
                };
                this.intersects = function(s, t, r) {
                    r = r || 0.001;
                    var q = this.getY(s),
                        m = this.getX(t);
                    return ((!isFinite(q) || Math.abs(t - q) < r) && (!isFinite(m) || Math.abs(s - m) < r))
                }
            }

            function d(r, s, m, p, o) {
                var n = r.line;
                var q = {},
                    t = 0,
                    u = 0;
                switch (true) {
                    case r.mode == l && r.labelPosition == "top":
                        u = o + r.labelYAdjust;
                        t = (s / 2) + r.labelXAdjust;
                        q.y = r.y1 + u;
                        q.x = (isFinite(n.m) ? n.getX(q.y) : r.x1) - t;
                        break;
                    case r.mode == l && r.labelPosition == "bottom":
                        u = m + o + r.labelYAdjust;
                        t = (s / 2) + r.labelXAdjust;
                        q.y = r.y2 - u;
                        q.x = (isFinite(n.m) ? n.getX(q.y) : r.x1) - t;
                        break;
                    case r.mode == i && r.labelPosition == "left":
                        t = p + r.labelXAdjust;
                        u = -(m / 2) + r.labelYAdjust;
                        q.x = r.x1 + t;
                        q.y = n.getY(q.x) + u;
                        break;
                    case r.mode == i && r.labelPosition == "right":
                        t = s + p + r.labelXAdjust;
                        u = -(m / 2) + r.labelYAdjust;
                        q.x = r.x2 - t;
                        q.y = n.getY(q.x) + u;
                        break;
                    default:
                        q.x = ((r.x1 + r.x2 - s) / 2) + r.labelXAdjust;
                        q.y = ((r.y1 + r.y2 - m) / 2) + r.labelYAdjust
                }
                return q
            }
            return j
        }
    }, {
        "../helpers.js": 5
    }]
}, {}, [6]);