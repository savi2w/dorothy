const Koa = require("koa");
const faker = require("faker");

const app = new Koa();

app.use(async (ctx) => {
  ctx.body = faker.commerce.productName();
});

app.listen(3114);
