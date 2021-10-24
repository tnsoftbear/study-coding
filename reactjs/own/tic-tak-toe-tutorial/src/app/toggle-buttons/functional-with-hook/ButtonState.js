export default class ButtonState
{
    constructor(isToggle, toggledCount)
    {
        console.log('ButtonState::constructor()');
        this.isToggle = isToggle ?? false;
        this.toggledCount = toggledCount ?? 0;
    }
}