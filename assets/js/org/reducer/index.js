import {combineReducers} from "redux";
import user from './user';
import org from './org';

export default combineReducers({
  user,
  org
});
