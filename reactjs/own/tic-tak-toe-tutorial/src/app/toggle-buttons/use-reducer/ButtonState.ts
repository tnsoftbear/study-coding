export interface ButtonStateInterface {
  isToggle: boolean;
  toggledCount: number;
}

export class ButtonState implements ButtonStateInterface {
  public isToggle: boolean;
  public toggledCount: number;
  constructor(isToggle?: boolean, toggledCount?: number) {
    this.isToggle = isToggle ?? false;
    this.toggledCount = toggledCount ?? 0;
  }
}

export interface ApplicationStateInterface {
  buttonStates: Array<ButtonStateInterface>;
}

export class ApplicationState implements ApplicationStateInterface {
  public buttonStates: Array<ButtonStateInterface>;
  constructor(length: number = 0) {
    this.buttonStates = Array.from({ length }, () => new ButtonState());
  }
}
