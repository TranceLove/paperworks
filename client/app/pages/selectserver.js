"use strict"

import React from "react";
import { Text, View, Card } from 'react-native';
import { Button } from 'react-native-elements'
import { Styles } from './../styles/styles'

let SelectServer = ({navigation}) => (
    <View style={Styles.SelectServer.content}>
        <View style={Styles.SelectServer.container}>
            <Button title="localhost" style={Styles.SelectServer.btnLocalhost} fontSize={20} />
            <Button title="Development" style={Styles.SelectServer.btnDevServer} fontSize={20} />
            <Button title="Staging" style={Styles.SelectServer.btnStagingServer} fontSize={20} />
            <Button title="Production" style={Styles.SelectServer.btnProductionServer} fontSize={20} />
        </View>
    </View>
)

SelectServer.navigationOptions = {}

export default SelectServer
