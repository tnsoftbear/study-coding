export interface ButtonStateInterface {
    isToggle: boolean;
    toggledCount: number;
}

export class ButtonState implements ButtonStateInterface
{
    public isToggle: boolean;
    public toggledCount: number;

    constructor(isToggle?: boolean, toggledCount?: number)
    {
        this.isToggle = isToggle ?? false;
        this.toggledCount = toggledCount ?? 0;
    }
}