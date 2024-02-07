import Controller from '@ember/controller';
import { action } from '@ember/object';
import ENV from 'frontend/config/environment';

export default class LoginController extends Controller {
    @action
    updateUsername(event) {
        this.set('username', event.target.value);
    }

    @action
    updatePassword(event) {
        this.set('password', event.target.value);
    }

    @action
    authenticate(event) {
        event.preventDefault();
        console.log('Authenticating...');
        console.log('Username:', this.username);
        let { username, password } = this;
        fetch(`${ENV.API_BASE_URL}/api/users/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => response.json())
            .then(data => {
                // Store the token in session or local storage
                // Example of storing in session storage:
                sessionStorage.setItem('token', data.token);
                console.log('Token stored successfully');
            })
            .catch(error => {
                console.error('Error occurred during login:', error);
            });

        console.log(`Logging in with username ${username} and password ${password}`);
    }
}
