import { useState, useEffect } from "react";

interface Props {
  seed: string;
}

interface UserInterface {
  gender: string;
  name: {
    title: string;
    first: string;
    last: string;
  };
  location: {
    street: {
      number: number;
      name: string;
    };
    city: string;
    state: string;
    country: string;
    postcode: number;
  };
}

interface UserResultsInterface {
  results: Array<UserInterface>;
}

const PersonPage = (props: Props) => {
  const [userData, setUserData] = useState<UserInterface>();

  useEffect(() => {
    fetchData(props.seed);
  }, [props.seed]);

  const fetchData = async (seed: string) => {
    try {
      const url = `https://randomuser.me/api/?seed=${seed}&noinfo`;
      const response = await fetch(url);
      if (!response.ok) {
        throw Error(response.statusText);
      }
      const data = (await response.json()) as UserResultsInterface;
      setUserData(data.results[0] ?? {});
    } catch (e) {
      console.log(e);
    }
  };

  const render = () => {
    if (userData) {
      const { title, first, last } = userData.name;
      const { street, city, state, country, postcode } = userData.location;
      return (
        <div>
          <div>
            {title} {first} {last}
          </div>
          <div>
            {street.name} {street.number}
          </div>
          <div>
            {city} {state} {country} {postcode}
          </div>
        </div>
      );
    }
    return "Loading...";
  };

  return <div>{render()}</div>;
};

export default PersonPage;
