const UrlPattern = require("url-pattern");

// eslint-disable-next-line no-unused-vars
function ll(...args) {
  global.console.log(JSON.stringify(args));
}

class MicroRouter {
  constructor() {
    // All registered handlers for get and post requests
    this.all = {
      get: [],
      post: []
    };
  }

  static init(req, res) {
    return {
      runI: -1,
      fns: [],
      req,
      res,
      err: null
    };
  }

  get(url, ...fns) {
    this.all.get[url] = (this.all.get[url] || []).concat(fns);
  }

  post(url, ...fns) {
    this.all.post[url] = (this.all.post[url] || []).concat(fns);
  }

  findRegisteredHandlers(ctx) {
    let found = [];
    const method = ctx.req.method.toLowerCase();
    const fns = this.all[method] || [];
    const { url } = ctx.req;
    for (const path in fns) {
      if (path === url) {
        found = found.concat(fns[path]);
      } else {
        const pattern = new UrlPattern(path);
        const matches = pattern.match(url);
        if (matches) {
          ctx.req.params = matches;
          found = found.concat(fns[path]);
        }
      }
    }
    return found;
  }

  static sendNotFound(req, res) {
    res.statusCode = 404;
    res.end();
  }

  find(ctx) {
    const fns = this.findRegisteredHandlers(ctx);
    if (fns.length) {
      return fns;
    }
    return [MicroRouter.sendNotFound];
  }

  runCurrent(ctx) {
    ctx.runI += 1;
    const current = ctx.fns[ctx.runI];
    if (typeof current !== "function") {
      return;
    }

    if (current.length < 3) {
      // sync way
      current(ctx.req, ctx.res);
      this.runCurrent(ctx);
    } else {
      // async way
      const next = err => {
        if (err) {
          ctx.err = err;
          // end routing chain, go to last handler
          // - it is next middleware registered after router
          ctx.runI = ctx.fns.length - 2;
        }
        this.runCurrent(ctx);
      };
      current(ctx.req, ctx.res, next);
    }
  }

  /**
   * Processing request according found route handlers
   * @param {http.ClientRequest} req
   * @param {http.ServerResponse} res
   * @param {Function} next probably next middleware
   */
  middleware() {
    return (req, res, next) => {
      const ctx = MicroRouter.init(req, res);
      ctx.fns = this.find(ctx);
      ctx.fns.push(() => {
        next(ctx.err);
      });
      try {
        this.runCurrent(ctx);
      } catch (err) {
        ctx.err = err;
        next(ctx.err);
      }
    };
  }
}

module.exports = () => {
  return new MicroRouter();
};
