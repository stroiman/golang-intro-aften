import * as uuid from 'uuid'

export const createMessage = () => ({
  id: uuid.v4()
});
