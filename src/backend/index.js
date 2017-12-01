import Application from 'koa';

const app = new Application();

app.use(async ctx => {
  ctx.body = 'Hello World';
});

app.listen(3000);
