import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllProfilesComponent } from './components/all-profiles/all-profiles.component';
import { JobsComponent } from './components/jobs/jobs.component';
import { LoginComponent } from './components/login/login.component';
import { PostsComponent } from './components/posts/posts.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ProfilesComponent } from './components/profiles/profiles.component';

const routes: Routes = [
  { path: 'jobs', component: JobsComponent },
  { path: 'profiles' , component: ProfilesComponent},
  { path: 'allprofiles',component: AllProfilesComponent},
  { path: 'profile/:id' , component: ProfileComponent},
  { path: 'login', component: LoginComponent},
  { path: 'posts', component: PostsComponent},
  { path: '', redirectTo: 'profiles', pathMatch: 'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
