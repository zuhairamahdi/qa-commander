import Route from '@ember/routing/route';

export default class LoginRoute extends Route {
  activate() {
    super.activate();
    document.body.classList.add('auth');
  }

  deactivate() {
    super.deactivate();
    document.body.classList.remove('auth');
  }
}