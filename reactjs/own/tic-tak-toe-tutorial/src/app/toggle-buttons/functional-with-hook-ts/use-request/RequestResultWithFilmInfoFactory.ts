import { FilmInfo } from "./FilmInfo";
import { RequestResultInterface } from "./RequestResult";
import RequestResultWithFilmInfo from "./RequestResultWithFilmInfo";

export interface FilmDataRequestInterface {
  title: string;
  episode_id: number | null;
}

export default class RequestResultWithFilmInfoFactory {
  static createByRequestResult(requestResult: RequestResultInterface): RequestResultWithFilmInfo 
  {
      const data = requestResult.data as FilmDataRequestInterface;
      return new RequestResultWithFilmInfo(requestResult, new FilmInfo(data.title, data.episode_id));
  }
}
