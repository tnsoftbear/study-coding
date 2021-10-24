export interface RequestResultInterface {
  data: {};
  isLoading: boolean;
  errorMessage: string;
  render: () => string;
}

export class RequestResult implements RequestResultInterface {
  public data: {};
  public isLoading: boolean;
  public errorMessage: string;

  constructor(
    data: {},
    isLoading: boolean = false,
    errorMessage: string = ''
  ) {
    this.data = data;
    this.isLoading = isLoading;
    this.errorMessage = errorMessage;
  }

  static constructorEmpty(): RequestResultInterface {
    return new RequestResult({});
  }

  static constructorLoading(): RequestResultInterface {
    return new RequestResult({}, true);
  }

  static constructorError(errorMessage: string): RequestResultInterface {
    return new RequestResult({}, false, errorMessage);
  }

  render(): string {
    if (this.isLoading) {
      return "Loading ...";
    }
    if (this.errorMessage) {
      return `Error: ${this.errorMessage}`;
    }
    return "";
  }
}
