"use strict"

import React from "react";
import { Text, View, Card } from 'react-native';
import { Button } from 'react-native-elements'
import { Styles } from './../styles/styles'
import Environment from './../services/Environment'
import { NavigationActions } from 'react-navigation'

class SelectServerView extends React.Component
{
    static navigationOptions = {}

    setServer(serverName:String)
    {
        Environment.set("serverName", serverName)
        this.props.navigation.dispatch(NavigationActions.reset({
            index: 0,
            actions: [
                NavigationActions.navigate({routeName:"Login"})
            ]
        }))
    }

    render()
    {
        return (
            <View style={Styles.SelectServer.content}>
                <View style={Styles.SelectServer.container}>
                    <Button title="localhost" style={Styles.SelectServer.btnLocalhost} onPress={()=>{this.setServer('localhost')}} fontSize={20} />
                    <Button title="Development" style={Styles.SelectServer.btnDevServer} onPress={()=>{this.setServer('development')}} fontSize={20} />
                    <Button title="Staging" style={Styles.SelectServer.btnStagingServer} onPress={()=>{this.setServer('staging')}} fontSize={20} />
                    <Button title="Production" style={Styles.SelectServer.btnProductionServer} onPress={()=>{this.setServer('production')}} fontSize={20} />
                </View>
            </View>
        )
    }
}

export default SelectServerView
