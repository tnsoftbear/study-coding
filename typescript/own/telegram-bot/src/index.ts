import { Telegraf } from "telegraf";
import { request as fetch } from "undici";
import ngrok from "ngrok";
import nconf from "nconf";

nconf.argv().env().file({ file: "config.json" });
console.log('port: ' + nconf.get("port"));

//Зашел на https://www.patreon.com/portal/registration/register-clients
//Прописал Redirect URIs: http://c85d-109-73-102-74.ngrok.io/oauth/patreon/callback

async function run() {
  await ngrok.authtoken(nconf.get("telegramApiKeyy"));
  const httpsHost = "http://c85d-109-73-102-74.ngrok.io"; // await ngrok.connect(5000);
  const bot = new Telegraf("5257675595:AAEy_kqPtl9tMsv8ypj-sEcByfQLh9JeCDI");
  bot.telegram.setWebhook(`${httpsHost}/secret-path`);
  // @ts-expect-error fixme
  bot.startWebhook("/secret-path", null, 5000);

  bot.command("patreon", (ctx) => {});

  bot.on("text", (ctx) => {
    ctx.telegram.sendMessage(ctx.message.chat.id, `Hello ${ctx.state.role}`);
    console.log(ctx.message.from);
    // Using context shortcut
    ctx.reply(`Hello ${ctx.state.role}`);
  });
}

run();
