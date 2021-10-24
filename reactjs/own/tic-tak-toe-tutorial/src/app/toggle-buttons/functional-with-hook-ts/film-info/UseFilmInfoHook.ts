import { useState, useEffect } from "react";
import {
  FilmInfoRequestResult,
  FilmInfoRequestResultInterface,
} from "./FilmInfo";

export const useFilmInfo = (id: number | null): FilmInfoRequestResultInterface => {
  const [filmInfoRequestResult, setFilmInfoRequestResult] =
    useState<FilmInfoRequestResultInterface>(
      FilmInfoRequestResult.constructorEmpty()
    );

  useEffect(() => {
    if (id !== null) {
      fetchFilmInfo(id + 1);
    }
  }, [id]);

  const fetchFilmInfo = async (idx: number) => {
    let cancelled = false;
    try {
      setFilmInfoRequestResult(FilmInfoRequestResult.constructorLoading());
      const response = await fetch(`https://swapi.dev/api/films/${idx}/`);
      if (!response.ok) {
        throw new Error(response.statusText);
      }
      const data = await response.json();
      if (!cancelled) {
        setFilmInfoRequestResult(
          new FilmInfoRequestResult({
            filmTitle: data.title,
            episodeId: data.episode_id,
          })
        );
      }
    } catch (e) {
      // Нет смысла обрабатывать ошибку, если загрузка была отменена
      if (!cancelled) {
        setFilmInfoRequestResult(
          FilmInfoRequestResult.constructorError("Error: " + e)
        );
      }
    }
    return () => {
      // Промис нельзя отменить, но можно проигнорировать результат таким способом
      cancelled = true;
    };
  };

  return filmInfoRequestResult;
};
