import Renderable from "./Renderable";

export interface FilmInfoInterface {
  filmTitle: string;
  episodeId: number | null;
  render: () => string;
}

export class FilmInfo implements FilmInfoInterface, Renderable {
  public filmTitle: string;
  public episodeId: number | null;

  constructor(title: string, episodeId: number | null) {
    this.filmTitle = title;
    this.episodeId = episodeId;
  }

  render(): string {
    if (this.episodeId) {
      return `${this.episodeId}: ${this.filmTitle}`;
    }
    return '';
  }
}
