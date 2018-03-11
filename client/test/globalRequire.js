const { JSDOM } = require('jsdom');

const jsdom = new JSDOM('<!doctype html><html><body></body></html>');
const { window } = jsdom;

const Enzyme = require('enzyme');
const Adapter = require('enzyme-adapter-react-16');
const chai = require('chai');
const sinonChai = require('sinon-chai');
chai.use(sinonChai);
chai.should();
global.expect = chai.expect;

Enzyme.configure({ adapter: new Adapter() });

function copyProps(src, target) {
  const props = Object.getOwnPropertyNames(src)
    .filter(prop => typeof target[prop] === 'undefined')
    .reduce((result, prop) => ({
      ...result,
      [prop]: Object.getOwnPropertyDescriptor(src, prop),
    }), {});
  Object.defineProperties(target, props);
}

global.fetch = require('node-fetch');
global.window = window;
global.document = window.document;
global.navigator = {
  userAgent: 'node.js',
};
copyProps(window, global);
