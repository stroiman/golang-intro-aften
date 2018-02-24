import sinon from 'sinon';

export const useSinon = () => {
  beforeEach(function() {
    this.sinon = sinon.createSandbox();
  });

  afterEach(function() {
    this.sinon.restore()
  });
}
