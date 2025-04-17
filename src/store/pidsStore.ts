import {defineStore} from 'pinia'

export const usePidsStore = defineStore('pids', {
    state: () => ({
        data: [
            {
                name: "pid_rotor1_rpm",
                Kp: 0.0001,
                Ki: 0.001,
                Kd: 0.0,
                integral_min: 0.0,
                integral_max: 1.0,
                inp_rise_deriative: 80,
                inp_fall_deriative: 60,
                min: 0.0,
                max: 1.0,
                preset_allowed_at_low: -1.0,
                preset_allowed_at_high: 1.0
            },
            {
                name: "pid_rotor2_rpm",
                Kp: 0.0001,
                Ki: 0.001,
                Kd: 0.0,
                integral_min: 0.0,
                integral_max: 1.0,
                inp_rise_deriative: 80,
                inp_fall_deriative: 60,
                min: 0.0,
                max: 1.0,
                preset_allowed_at_low: -1.0,
                preset_allowed_at_high: 1.0
            },
            {
                name: "pid_rotor3_rpm",
                Kp: 0.0001,
                Ki: 0.001,
                Kd: 0.0,
                integral_min: 0.0,
                integral_max: 1.0,
                inp_rise_deriative: 80,
                inp_fall_deriative: 60,
                min: 0.0,
                max: 1.0,
                preset_allowed_at_low: -1.0,
                preset_allowed_at_high: 1.0
            },
            {
                name: "pid_rotor4_rpm",
                Kp: 0.0001,
                Ki: 0.001,
                Kd: 0.0,
                integral_min: 0.0,
                integral_max: 1.0,
                inp_rise_deriative: 80,
                inp_fall_deriative: 60,
                min: 0.0,
                max: 1.0,
                preset_allowed_at_low: -1.0,
                preset_allowed_at_high: 1.0
            },
            {
                name: "pid_rotor5_rpm",
                Kp: 0.0001,
                Ki: 0.001,
                Kd: 0.0,
                integral_min: 0.0,
                integral_max: 1.0,
                inp_rise_deriative: 80,
                inp_fall_deriative: 60,
                min: 0.0,
                max: 1.0,
                preset_allowed_at_low: -1.0,
                preset_allowed_at_high: 1.0
            }
        ]
    }),
})
