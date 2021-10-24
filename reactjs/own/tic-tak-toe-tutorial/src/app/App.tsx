import { Component } from "react";
import GameBoard from "./game/game-js/GameBoard.js";
import GameBoardTs from "./game/game-ts/GameBoard";
import GameBoardFunctional from "./game/game-functional-ts/GameBoardFunctional";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import ScoreBoard from "./score-board/ScoreBoard";
import ToggleButtonsPage from "./toggle-buttons/class-component/ToggleButtonsPage.js";
import ToggleButtonsContextPage from "./toggle-buttons/context-api/ToggleButtonsContextPage.js";
import ToggleButtonsFunctionalPage from "./toggle-buttons/functional/ToggleButtonsFunctionalPage.js";
import ToggleButtonsFunctionalWithHooksPage from "./toggle-buttons/functional-with-hook/ToggleButtonsFunctionalWithHooksPage.js";
import ToggleButtonsFunctionalWithHooksAndTsPage from "./toggle-buttons/functional-with-hook-ts/ToggleButtonsFunctionalWithHooksAndTsPage";
//import ToggleButtonsMobxPage from "./toggle-buttons/mobx/ToggleButtonsMobxPage.js";
import ToggleButtonsUseReducerPage from "./toggle-buttons/use-reducer/ToggleButtonsUseReducerPage";

export default class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <nav>
            <ul>
              <li>
              Game class-components <Link to="/game-js">(JS)</Link> | <Link to="/game-ts">(TS)</Link>
              </li>
              <li>
              Game functional-components <Link to="/game-functional-ts">(TS)</Link>
              </li>
              <li>
                <Link to="/score-board">Score</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-class-component">Toggle buttons (class component)</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-context">Toggle buttons (context)</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-functional">Toggle buttons (functional)</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-functional-with-hooks">Toggle buttons (functional with hook)</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-functional-with-hooks-and-ts">Toggle buttons (functional with hook and typescript)</Link>
              </li>
              <li>
                <Link to="/toggle-buttons-use-reducer">Toggle buttons (use reducer)</Link>
              </li>
              {/* <li>
                <Link to="/toggle-buttons-mobx">Toggle buttons (mobx)</Link>
              </li> */}
            </ul>
          </nav>
          <Switch>
            <Route path="/game-js">
              <GameBoard />
            </Route>
            <Route path="/game-ts">
              <GameBoardTs />
            </Route>
            <Route path="/game-functional-ts">
              <GameBoardFunctional />
            </Route>
            <Route path="/score-board">
              <ScoreBoard />
            </Route>
            <Route path="/toggle-buttons-class-component">
              <ToggleButtonsPage />
            </Route>
            <Route path="/toggle-buttons-context">
              <ToggleButtonsContextPage count="5" />
            </Route>
            <Route path="/toggle-buttons-functional">
              <ToggleButtonsFunctionalPage count="5" />
            </Route>
            <Route path="/toggle-buttons-functional-with-hooks">
              <ToggleButtonsFunctionalWithHooksPage count="5" />
            </Route>
            <Route path="/toggle-buttons-functional-with-hooks-and-ts">
              <ToggleButtonsFunctionalWithHooksAndTsPage count={5} />
            </Route>
            <Route path="/toggle-buttons-use-reducer">
              <ToggleButtonsUseReducerPage count={3} />
            </Route>
            {/* <Route path="/toggle-buttons-mobx">
              <ToggleButtonsMobxPage />
            </Route> */}
          </Switch>
        </div>
      </Router>
    );
  }
}
