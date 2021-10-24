import { RequestResult } from "./RequestResult";
import { FilmInfoInterface } from "./FilmInfo";
import Renderable from "./Renderable";

export default class RequestResultWithFilmInfo implements Renderable
{
    public requestResult: RequestResult;
    public filmInfo: FilmInfoInterface;
    constructor(requestResult: RequestResult, filmInfo: FilmInfoInterface)
    {
        this.requestResult = requestResult;
        this.filmInfo = filmInfo;
    }

    render(): string
    {
        let output = this.requestResult.render();
        if (output !== '') {
            return output;
        }
        return this.filmInfo.render();
    }
}