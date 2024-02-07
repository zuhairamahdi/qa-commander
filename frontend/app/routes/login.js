import Route from '@ember/routing/route';
import ENV from 'frontend/config/environment';

export default class LoginRoute extends Route {
    model() {
        return {
            username: '',
            password: ''
        };
        }

    setupController(controller, model) {
        controller.setProperties(model);
    }
    
    activate() {
        super.activate();
        console.log(ENV.API_BASE_URL);
        document.body.classList.add('auth');
    }

    deactivate() {
        super.deactivate();
        document.body.classList.remove('auth');
    }
}