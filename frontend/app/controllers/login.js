import Controller from '@ember/controller';
import { action } from '@ember/object';

export default class LoginController extends Controller {
    @action
    authenticate(event) {
        event.preventDefault();
        let { username, password } = this;
        console.log(`Logging in with username ${username} and password ${password}`);
    }
}
