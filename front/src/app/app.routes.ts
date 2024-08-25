import { Routes } from '@angular/router';
import {PageNotFoundComponent} from './page-not-found/page-not-found.component'
import {HomeComponent} from './home/home.component'
import {ExternoComponent} from './externo/externo.component'
import {LoginComponent} from './login/login.component'
import {TesteComponent} from './teste/teste.component'

export const routes: Routes = [
    { path: '', component: HomeComponent},
    {
        path: 'externo', 
        component: ExternoComponent,
        children: [
            { path: 'login', component: LoginComponent},
        ]
    },
    { path: 'teste', component: TesteComponent},
    { path: '**', component: PageNotFoundComponent},
];
