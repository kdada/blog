import Application from 'koa';
import Router from 'koa-router';
import serve from 'koa-static';

function start() {
  console.log(process.cwd());
  const app = new Application();
  const router = new Router();
  router.get('/', function(ctx) {
    ctx.response.body = 'test';
  });
  app.use(router.routes()).use(router.allowedMethods());
  app.use(serve('./web/pages/'));
  app.listen(3000);
}

start();
