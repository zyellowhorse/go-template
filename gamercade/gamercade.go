package gamercade

import (
	"errors"
	"math"
)

// Audio Api
func PlayBgm(bgm_index uint) {
	Play_bgm(int32(bgm_index))
}

func PlaySfx(sfx_index uint, channel uint) {
	Play_sfx(int32(sfx_index), int32(channel))
}

func StopBgm() {
	Stop_bgm()
}

func StopChannel(channel uint) {
	Stop_channel(int32(channel))
}

func PlayNote(note_id uint, instrument_index uint, channel uint) {
	Play_note(int32(note_id), int32(instrument_index), int32(channel))
}

func PlayFrequency(frequency float32, instrument_index uint, channel uint) {
	Play_frequency(frequency, int32(instrument_index), int32(channel))
}

// Data Api
func Height_screen() uint {
	return uint(Height())
}

func Width_screen() uint {
	return uint(Width())
}

func Fps_screen() uint {
	return uint(Fps())
}

func FrameTime() float32 {
	return Frame_time()
}

func SpriteSheetCount() uint {
	return uint(Sprite_sheet_count())
}

func PaletteCount() uint {
	return uint(Palette_count())
}

func SpriteHeight(sprite_sheet uint) uint {
	return uint(Sprite_height(int32(sprite_sheet)))
}

func SpriteWidth(sprite_sheet uint) uint {
	return uint(Sprite_width(int32(sprite_sheet)))
}

func SpriteCount(sprite_sheet uint) uint {
	return uint(Sprite_count(int32(sprite_sheet)))
}

func BgmLengthSecs(bgm_index uint) float32 {
	return Bgm_length_secs(int32(bgm_index))
}

func BgmLengthFrames(bgm_index uint) uint {
	return uint(Bgm_length_frames(int32(bgm_index)))
}

func SfxLengthSecs(sfx_index uint) float32 {
	return Sfx_length_secs(int32(sfx_index))
}

func SfxLengthFrames(sfx_index uint) uint {
	return uint(Sfx_length_frames(int32(sfx_index)))
}

// Graphics Params Api
func PaletteIndex(palette_index uint) int32 {
	return Palette_index(int32(palette_index))
}

func SpriteSheetIndex(sprite_sheet_index uint) int32 {
	return Sprite_sheet_index(int32(sprite_sheet_index))
}

func SpriteIndex(sprite_index uint) int32 {
	return Sprite_index(int32(sprite_index))
}

func ColorIndex(color_index uint) int32 {
	return Color_index(int32(color_index))
}

func FlipX(flip_x bool) int32 {
	return Flip_x(boolToInt32(flip_x))
}

func FlipY(flip_y bool) int32 {
	return Flip_y(boolToInt32(flip_y))
}

func GraphicsParameters(
	palette_index uint,
	sprite_sheet_index uint,
	sprite_index uint,
	color_index uint,
	flip_x bool,
	flip_y bool,
) int32 {
	return Graphics_parameters(int32(palette_index), int32(sprite_sheet_index), int32(sprite_index), int32(color_index), boolToInt32(flip_x), boolToInt32(flip_y))
}

// Draw Api
func ClearScreen(graphics_parameters int32) {
	Clear_screen(graphics_parameters)
}

func SetPixel(graphics_parameters int32, x int32, y int32) {
	Set_pixel(graphics_parameters, x, y)
}

func DrawCircle(graphics_parameters int32, x int32, y int32, radius int32) {
	Circle(graphics_parameters, x, y, radius)
}

func DrawCircleFilled(graphics_parameters int32, x int32, y int32, radius int32) {
	Circle_filled(graphics_parameters, x, y, radius)
}

func DrawRect(graphics_parameters int32, x int32, y int32, rect_width int32, rect_height int32) {
	Rect(graphics_parameters, x, y, rect_width, rect_height)
}

func RectFilled(graphics_parameters int32, x int32, y int32, rect_width int32, rect_height int32) {
	Rect_filled(graphics_parameters, x, y, rect_width, rect_height)
}

func DrawLine(graphics_parameters int32, x0 int32, y0 int32, x1 int32, y1 int32) {
	Line(graphics_parameters, x0, y0, x1, y1)
}

func DrawSprite(graphics_parameters int32, transparency_mask int64, x int32, y int32) {
	Sprite(graphics_parameters, transparency_mask, x, y)
}

// Text Api
func ConsoleLog(text string) {
	Console_log([]uint8(text), uint(len(text)))
}

// Random Api
func SetSeed(seed int32) {
	Set_seed(seed)
}

func RandomIntRange(min int32, max int32) int32 {
	return Random_int_range(min, max)
}

func RandomFloat() float32 {
	return Random_float()
}

func RandomFloatRange(min float32, max float32) float32 {
	return Random_float_range(min, max)
}

// Input Api
func ButtonAPressed(player_id uint) bool {
	return int32ToBool(Button_a_pressed(int32(player_id)))
}

func ButtonAReleased(player_id uint) bool {
	return int32ToBool(Button_a_released(int32(player_id)))
}

func ButtonAHeld(player_id uint) bool {
	return int32ToBool(Button_a_held(int32(player_id)))
}

func ButtonBPressed(player_id uint) bool {
	return int32ToBool(Button_b_pressed(int32(player_id)))
}

func ButtonBReleased(player_id uint) bool {
	return int32ToBool(Button_b_released(int32(player_id)))
}

func ButtonBHeld(player_id uint) bool {
	return int32ToBool(Button_b_held(int32(player_id)))
}

func ButtonCPressed(player_id uint) bool {
	return int32ToBool(Button_c_pressed(int32(player_id)))
}

func ButtonCReleased(player_id uint) bool {
	return int32ToBool(Button_c_released(int32(player_id)))
}

func ButtonCHeld(player_id uint) bool {
	return int32ToBool(Button_c_held(int32(player_id)))
}

func ButtonDPressed(player_id uint) bool {
	return int32ToBool(Button_d_pressed(int32(player_id)))
}

func ButtonDReleased(player_id uint) bool {
	return int32ToBool(Button_d_released(int32(player_id)))
}

func ButtonDHeld(player_id uint) bool {
	return int32ToBool(Button_d_held(int32(player_id)))
}

func ButtonUpPressed(player_id uint) bool {
	return int32ToBool(Button_up_pressed(int32(player_id)))
}

func ButtonUpReleased(player_id uint) bool {
	return int32ToBool(Button_up_released(int32(player_id)))
}

func ButtonUpHeld(player_id uint) bool {
	return int32ToBool(Button_up_held(int32(player_id)))
}

func ButtonDownPressed(player_id uint) bool {
	return int32ToBool(Button_down_pressed(int32(player_id)))
}

func ButtonDownReleased(player_id uint) bool {
	return int32ToBool(Button_down_released(int32(player_id)))
}

func ButtonDownHeld(player_id uint) bool {
	return int32ToBool(Button_down_held(int32(player_id)))
}

func ButtonLeftPressed(player_id uint) bool {
	return int32ToBool(Button_left_pressed(int32(player_id)))
}

func ButtonLeftReleased(player_id uint) bool {
	return int32ToBool(Button_left_released(int32(player_id)))
}

func ButtonLeftHeld(player_id uint) bool {
	return int32ToBool(Button_left_held(int32(player_id)))
}

func ButtonRightPressed(player_id uint) bool {
	return int32ToBool(Button_right_pressed(int32(player_id)))
}

func ButtonRightReleased(player_id uint) bool {
	return int32ToBool(Button_right_released(int32(player_id)))
}

func ButtonRightHeld(player_id uint) bool {
	return int32ToBool(Button_right_held(int32(player_id)))
}

func ButtonStartPressed(player_id uint) bool {
	return int32ToBool(Button_start_pressed(int32(player_id)))
}

func ButtonStartReleased(player_id uint) bool {
	return int32ToBool(Button_start_released(int32(player_id)))
}

func ButtonStartHeld(player_id uint) bool {
	return int32ToBool(Button_start_held(int32(player_id)))
}

func ButtonSelectPressed(player_id uint) bool {
	return int32ToBool(Button_select_pressed(int32(player_id)))
}

func ButtonSelectReleased(player_id uint) bool {
	return int32ToBool(Button_select_released(int32(player_id)))
}

func ButtonSelectHeld(player_id uint) bool {
	return int32ToBool(Button_select_held(int32(player_id)))
}

func ButtonLeftShoulderPressed(player_id uint) bool {
	return int32ToBool(Button_left_shoulder_pressed(int32(player_id)))
}

func ButtonLeftShoulderReleased(player_id uint) bool {
	return int32ToBool(Button_left_shoulder_released(int32(player_id)))
}

func ButtonLeft_shoulder_held(player_id uint) bool {
	return int32ToBool(Button_left_shoulder_held(int32(player_id)))
}

func ButtonRightShoulderPressed(player_id uint) bool {
	return int32ToBool(Button_right_shoulder_pressed(int32(player_id)))
}

func ButtonRightShoulderReleased(player_id uint) bool {
	return int32ToBool(Button_right_shoulder_released(int32(player_id)))
}

func ButtonRight_shoulder_held(player_id uint) bool {
	return int32ToBool(Button_right_shoulder_held(int32(player_id)))
}

func ButtonLeftStickPressed(player_id uint) bool {
	return int32ToBool(Button_left_stick_pressed(int32(player_id)))
}

func ButtonLeftStickReleased(player_id uint) bool {
	return int32ToBool(Button_left_stick_released(int32(player_id)))
}

func ButtonLeftStickHeld(player_id uint) bool {
	return int32ToBool(Button_left_stick_held(int32(player_id)))
}

func ButtonRightStickPressed(player_id uint) bool {
	return int32ToBool(Button_right_stick_pressed(int32(player_id)))
}

func ButtonRightStickReleased(player_id uint) bool {
	return int32ToBool(Button_right_stick_released(int32(player_id)))
}

func ButtonRightStickHeld(player_id uint) bool {
	return int32ToBool(Button_right_stick_held(int32(player_id)))
}

func ButtonLeftTriggerPressed(player_id uint) bool {
	return int32ToBool(Button_left_trigger_pressed(int32(player_id)))
}

func ButtonLeftTriggerReleased(player_id uint) bool {
	return int32ToBool(Button_left_trigger_released(int32(player_id)))
}

func ButtonLeftTriggerHeld(player_id uint) bool {
	return int32ToBool(Button_left_trigger_held(int32(player_id)))
}

func ButtonRightRriggerPressed(player_id uint) bool {
	return int32ToBool(Button_right_trigger_pressed(int32(player_id)))
}

func ButtonRightTriggerReleased(player_id uint) bool {
	return int32ToBool(Button_right_trigger_released(int32(player_id)))
}

func ButtonRightTriggerHeld(player_id uint) bool {
	return int32ToBool(Button_right_trigger_held(int32(player_id)))
}

func AnalogLeftX(player_id uint) (float32, error) {
	return float32ToOption(Analog_left_x(int32(player_id)))
}

func AnalogLeftY(player_id uint) (float32, error) {
	return float32ToOption(Analog_left_y(int32(player_id)))
}

func AnalogRightX(player_id uint) (float32, error) {
	return float32ToOption(Analog_right_x(int32(player_id)))
}

func AnalogRightY(player_id uint) (float32, error) {
	return float32ToOption(Analog_right_y(int32(player_id)))
}

func TriggerLeft(player_id uint) (float32, error) {
	return float32ToOption(Trigger_left(int32(player_id)))
}

func TriggerRight(player_id uint) (float32, error) {
	return float32ToOption(Trigger_right(int32(player_id)))
}

func RawInputState(player_id uint) (int64, error) {
	val := Raw_input_state(int32(player_id))
	if val < 0 {
		return 0, errors.New("unable to get Raw Input State")
	} else {
		return val, nil
	}
}

// Multiplayer Api
func NumPlayers() uint {
	return uint(Num_players())
}

func IsLocalPlayer(player_id uint) bool {
	return int32ToBool(Is_local_player(int32(player_id)))
}

func IsRemotePlayer(player_id uint) bool {
	return int32ToBool(Is_remote_player(int32(player_id)))
}

func float32ToOption(value float32) (float32, error) {
	if !math.IsInf(float64(value), 0) && !math.IsNaN(float64(value)) {
		return value, nil
	} else {
		return 0, errors.New("value is not a number")
	}
}

func int32ToBool(value int32) bool {
	switch value {
	case 0:
		return false
	case 1:
		return true
	default:
		return false
	}
}

func boolToInt32(value bool) int32 {
	switch value {
	case false:
		return 0
	case true:
		return 1
	default:
		return 0
	}
}
