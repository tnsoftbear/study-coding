// eslint-disable-next-line no-unused-vars
function ll(...args) {
  global.console.log(JSON.stringify(args));
}

class MicroExpress {
  constructor() {
    this.middlewares = [];
  }

  static runFinal() {}

  findByIndex(i) {
    let fn = this.middlewares[i];
    if (typeof fn !== "function") {
      fn = MicroExpress.runFinal;
    }
    return fn;
  }

  detectIndex(ctx) {
    let index;
    for (index = ctx.runI; index < this.middlewares.length; index += 1) {
      const fn = this.findByIndex(index);
      if (ctx.err && fn.length === 4) {
        return index;
      }
      if (!ctx.err && fn.length < 4) {
        return index;
      }
    }
    return index;
  }

  runCurrent(ctx) {
    ctx.runI += 1;
    ctx.runI = this.detectIndex(ctx);
    const current = this.findByIndex(ctx.runI);
    const next = err => {
      if (err) {
        ctx.err = err;
      }
      this.runCurrent(ctx);
    };
    try {
      if (ctx.err) {
        current(ctx.err, ctx.req, ctx.res, next);
      } else {
        current(ctx.req, ctx.res, next);
      }
    } catch (err) {
      ctx.err = err;
      this.runCurrent(ctx);
    }
  }

  static init(req, res) {
    return {
      runI: -1,
      req,
      res,
      err: null
    };
  }

  handler() {
    return (req, res) => {
      this.middlewares.push(MicroExpress.runFinal);
      const ctx = MicroExpress.init(req, res);
      try {
        this.runCurrent(ctx);
      } catch (err) {
        ctx.err = err;
        this.runCurrent(ctx);
      }
    };
  }

  use(middleware) {
    this.middlewares.push(middleware);
  }
}

module.exports = () => {
  return new MicroExpress();
};
