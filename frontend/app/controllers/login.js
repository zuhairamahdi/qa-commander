import Controller from '@ember/controller';
import { action } from '@ember/object';
import RouterService from '@ember/routing/router-service';
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
        // if token is already present, redirect to home
        if (sessionStorage.getItem('token')) {
            RouterService.transitionTo('dashboard');
            return;
        }
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

                sessionStorage.setItem('token', data.token);
                //add isAuthenticated to session
                this.session.set('isAuthenticated', true);
            })
            .catch(error => {
                console.error('Error occurred during login:', error);
            });

        console.log(`Logging in with username ${username} and password ${password}`);
    }
}
