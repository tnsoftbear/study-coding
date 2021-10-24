export default class ButtonState
{
    constructor(isToggle, toggledCount)
    {
        this.isToggle = isToggle ?? false;
        this.toggledCount = toggledCount ?? 0;
    }
}