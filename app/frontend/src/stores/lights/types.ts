export interface LightState {
  color: number
  speed: number
  brightness: number
  mode: number
  enabled: boolean
}

export interface LightMode {
  code: number
  features: number
  name: string
}

export interface EffectParams {
  Color: number
  Speed: number
  Brightness: number
}

export type LightSetter = (
  mode: number,
  colorIndex: number,
  brightness: number,
  speed: number
) => Promise<void>

export interface Color {
  R: number
  G: number
  B: number
}

type RandomColor = 'random'

export type PreviewColor = Color | RandomColor | undefined
