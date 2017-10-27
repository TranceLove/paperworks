"use strict"

import React from "react";
import { Text, View, Card } from 'react-native';
import { Button } from 'react-native-elements'
import { Styles } from './../styles/styles'
import Environment from './../services/Environment'
import { NavigationActions } from 'react-navigation'

function setServer(navigation,serverName)
{
    Environment.set("serverName", serverName)
    navigation.dispatch(NavigationActions.reset({
        index: 0,
        actions: [
            NavigationActions.navigate({routeName:"Login"})
        ]
    }))
}

let SelectServer = ({navigation}) => (
    <View style={Styles.SelectServer.content}>
        <View style={Styles.SelectServer.container}>
            <Button title="localhost" style={Styles.SelectServer.btnLocalhost} onPress={()=>{setServer(navigation,'localhost')}} fontSize={20} />
            <Button title="Development" style={Styles.SelectServer.btnDevServer} onPress={()=>{setServer(navigation,'development')}} fontSize={20} />
            <Button title="Staging" style={Styles.SelectServer.btnStagingServer} onPress={()=>{setServer(navigation,'staging')}} fontSize={20} />
            <Button title="Production" style={Styles.SelectServer.btnProductionServer} onPress={()=>{setServer(navigation,'production')}} fontSize={20} />
        </View>
    </View>
)

SelectServer.navigationOptions = {}

export default SelectServer
