// export const A_TOGGLE = 'toggle';
// export const A_DROP_RANDOM = 'drop-random';
// export const A_ADD_TOGGLE_BUTTON = 'add-toggle-button';
// export const A_REMOVE_TOGGLE_BUTTON = 'remove-toggle-button';
// export const A_UNDO = 'undo';
// export const A_REDO = 'redo';

export type ActionType = 
  | { type: 'toggle', buttonIndex: number, newIsToggle: boolean }
  | { type: 'drop-random', droppedIndex: number }
  | { type: 'add-toggle-button' }
  | { type: 'remove-toggle-button', removeIndex: number }
  | { type: 'undo' }
  | { type: 'redo' }