import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { JwtHelperService, JWT_OPTIONS } from '@auth0/angular-jwt';

import { TokenInterceptor } from './interceptor/token-interceptor';
import { AuthService } from './services/auth/auth.service';
import { UserService } from './services/user/user.service';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { HeaderComponent } from './components/header/header.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { LoginPageComponent } from './components/login-page/login-page.component';
import { UserProfileComponent } from './components/user-profile/user-profile.component';
import { JobComponent } from './components/job/job.component';
import { AddJobComponent } from './components/add-job/add-job.component';
import { CommentComponent } from './components/comment/comment.component';
import { AddCommentComponent } from './components/add-comment/add-comment.component';
import { WageComponent } from './components/wage/wage.component';
import { AddWageComponent } from './components/add-wage/add-wage.component';
import { InterviewComponent } from './components/interview/interview.component';
import { AddInterviewComponent } from './components/add-interview/add-interview.component';


import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatRadioModule } from '@angular/material/radio';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatListModule } from '@angular/material/list';
import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { CompaniesComponent } from './components/companies/companies.component';
import { RegisterCompanyRequestsComponent } from './components/register-company-requests/register-company-requests.component';
import { CompanyComponent } from './components/company/company.component';

const MaterialComponents = [
  MatToolbarModule,
  MatButtonModule,
  MatCardModule,
  MatFormFieldModule,
  MatInputModule,
  MatRadioModule,
  MatProgressSpinnerModule,
  MatSnackBarModule,
  MatListModule
]

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    RegistrationPageComponent,
    LoginPageComponent,
    UserProfileComponent,
    JobComponent,
    CommentComponent,
    WageComponent,
    InterviewComponent,
    AddJobComponent,
    AddCommentComponent,
    AddWageComponent,
    AddInterviewComponent,
    RegisterCompanyComponent,
    CompaniesComponent,
    RegisterCompanyRequestsComponent,
    CompanyComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialComponents
  ],
  providers: [
    {
      provide: JWT_OPTIONS,
      useValue: JWT_OPTIONS
    },
    JwtHelperService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true
    },
    AuthService,
    UserService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
