"use strict"

import React from "react";
import { Text, View, Card } from 'react-native';
import { Button } from 'react-native-elements'
import { Styles } from './../styles/styles'
import { NavigationActions } from 'react-navigation'

class LoginView extends React.Component
{
    doLogin(method:String)
    {
        this.props.navigation.dispatch(NavigationActions.reset({
            index: 0,
            actions: [
                NavigationActions.navigate({routeName:"MainView"})
            ]
        }))
    }

    render()
    {
        return (
            <View style={Styles.Login.content}>
                <View style={Styles.Login.container}>
                    <Button
                        title="Login with Twitter"
                        icon={{name: 'twitter', type: 'entypo'}}
                        style={Styles.Login.twitterLogin}
                        backgroundColor="#1dcaff"
                        fontSize={20}
                    />
                    <Button
                        title="Login with Facebook"
                        icon={{name: 'facebook', type: 'entypo'}}
                        style={Styles.Login.facebookLogin}
                        backgroundColor="#3b5998"
                        fontSize={20}
                    />
                    <Button
                        title="Continue without login"
                        icon={{name: 'map', type: 'entypo'}}
                        fontSize={19}
                        onPress={() => this.doLogin("nologin")}
                    />
                </View>
            </View>
        )
    }
}

export default LoginView
