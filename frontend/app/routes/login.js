import Route from '@ember/routing/route';
import { inject as service } from '@ember/service';


export default class LoginRoute extends Route {
    @service router;
    beforeModel() {
        console.log('beforeModel');
        //check if token is present in session or local storage and redirect to dashboard
        if (sessionStorage.getItem('token')) {
            //redirect to dashboard
            this.router.transitionTo('dashboard');
            return;
        }
      }
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
        // if token is already present, redirect to dashboard and dont add auth class to body
        document.body.classList.add('auth');
    }

    deactivate() {
        super.deactivate();
        document.body.classList.remove('auth');
    }
}