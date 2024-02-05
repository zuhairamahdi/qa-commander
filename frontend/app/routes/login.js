import Route from '@ember/routing/route';
import { action } from '@ember/object';
import RouteClassMixin from 'frontend/mixins/route-class';

export default class LoginRoute extends Route.extend(RouteClassMixin) {
  @action
  authenticate(event) {
    event.preventDefault();

    let { username, password } = this.controller;

    
    // TODO: Authenticate the user with the server
    console.log(`Logging in with username ${username} and password ${password}`);
  }
  activate() {
    super.activate();
    document.body.classList.add('auth');
  }

  deactivate() {
    super.deactivate();
    document.body.classList.remove('auth');
  }
}