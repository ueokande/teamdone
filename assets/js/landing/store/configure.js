import logger  from 'redux-logger';
import thunk from 'redux-thunk';
import { createStore, applyMiddleware } from 'redux';
import rootReducer from '../reducer/index';

export default function createStoreWithMiddleware() {
  const store = applyMiddleware(
    logger,
    thunk
  )(createStore);
  return store(rootReducer);
}
