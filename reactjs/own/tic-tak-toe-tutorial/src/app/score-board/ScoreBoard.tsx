import { Component } from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Clock from "../common/clock/Clock.js";
import PersonPage from "./PersonPage";

interface PersonInterface {
  login: { uuid: string };
  name: { first: string; last: string };
  email: string;
  phone: string;
  picture: { thumbnail: string };
  seed: string;
}

class ScoreBoard extends Component<{}, any> {
  constructor(props: Array<any>) {
    super(props);
    this.state = {
      data: [],
    };
  }

  componentDidMount() {
    this.fetchData();
  }

  // fetchData = () => {
  //   return fetch(
  //     "https://randomuser.me/api/?results=5&noinfo&inc=name,email,phone,login,picture"
  //   )
  //     .then((response) => response.json())
  //     .then((data) => {
  //       this.setState({ data: data.results });
  //     });
  // };

  async fetchData() {
    try {
      const response = await fetch(
        "https://randomuser.me/api/?results=5"
      );
      if (!response.ok) {
        throw Error(response.statusText);
      }
      const data = await response.json();
      this.setState({ data: data.results });
    } catch (e) {
      console.log(e);
    }
  }

  render() {
    return (
      <div>
        <h1>Scores</h1>
        <Router>
          <table>
            {this.state.data &&
              this.state.data.map((person: PersonInterface) => {
                const to = `/person/${person.login.uuid}`;
                return (
                  <tbody key={person.login.uuid}>
                    <tr>
                      <td>
                        <Link to={to}>
                          {person.name.first}, {person.name.last}
                        </Link>
                      </td>
                      <td>{person.email}</td>
                      <td>{person.phone}</td>
                      <td>
                        <img
                          src={person.picture?.thumbnail}
                          alt={person.name.last}
                        />
                      </td>
                    </tr>
                  </tbody>
                );
              })}
          </table>
          <Switch>
            <Route
              path="/person/:seed"
              render={({ match }) => {
                const { seed } = match.params;
                return <PersonPage seed={seed} />;
              }}
            />
          </Switch>
        </Router>
        <div>
          <Clock textColor="orange" />
        </div>
      </div>
    );
  }
}

export default ScoreBoard;
