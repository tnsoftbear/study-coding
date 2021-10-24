import GameMain from "./game-main.js";

console.log("Starting");

function main() {
  const gameMain = new GameMain();
  const img = document.createElement("img");
  img.src = gameMain.imageSrc();
  document.body.appendChild(img);
}

window.addEventListener('DOMContentLoaded', (event) => {
    main();
});

console.log("Ending");
