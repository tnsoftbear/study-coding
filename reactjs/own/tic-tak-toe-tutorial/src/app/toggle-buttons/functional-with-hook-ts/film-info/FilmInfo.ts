export interface FilmInfoInterface {
  filmTitle: string;
  episodeId: number | null;
}

export interface FilmInfoRequestResultInterface {
  filmInfo: FilmInfoInterface;
  isLoading: boolean;
  errorMessage: string;
  render: () => string;
}

export class FilmInfoRequestResult implements FilmInfoRequestResultInterface {
  public filmInfo: FilmInfoInterface;
  public isLoading: boolean;
  public errorMessage: string;

  constructor(
    filmInfo: FilmInfoInterface,
    isLoading: boolean = false,
    errorMessage: string = ""
  ) {
    this.filmInfo = filmInfo;
    this.isLoading = isLoading;
    this.errorMessage = errorMessage;
  }

  static constructorEmpty(): FilmInfoRequestResultInterface {
    return new FilmInfoRequestResult({ filmTitle: "", episodeId: null });
  }

  static constructorLoading(): FilmInfoRequestResultInterface {
    return new FilmInfoRequestResult({ filmTitle: "", episodeId: null }, true);
  }

  static constructorError(
    errorMessage: string
  ): FilmInfoRequestResultInterface {
    return new FilmInfoRequestResult(
      { filmTitle: "", episodeId: null },
      false,
      errorMessage
    );
  }

  render(): string {
    if (this.isLoading) {
      return "Loading ...";
    }
    if (this.errorMessage) {
      return `Error: ${this.errorMessage}`;
    }
    if (this.filmInfo.episodeId) {
      return `${this.filmInfo.episodeId}: ${this.filmInfo.filmTitle}`;
    }
    return "";
  }
}
