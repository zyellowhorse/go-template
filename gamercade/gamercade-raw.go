package gamercade

// Raw bindings directly from WASM.

// Audio
func Play_bgm(bgm_index int32)
func Play_sfx(sfx_index int32, channel int32)
func Stop_bgm()
func Stop_channel(channel int32)
func Play_note(note_id int32, instrument_index int32, channel int32)
func Play_frequency(frequency float32, instrument_index int32, channel int32)

// Data
func Height() int32
func Width() int32
func Fps() int32
func Frame_time() float32
func Sprite_sheet_count() int32
func Palette_count() int32
func Sprite_height(sprite_sheet int32) int32
func Sprite_width(sprite_sheet int32) int32
func Sprite_count(sprite_sheet int32) int32
func Bgm_length_secs(bgm_index int32) float32
func Bgm_length_frames(bgm_index int32) int32
func Sfx_length_secs(sfx_index int32) float32
func Sfx_length_frames(sfx_index int32) int32

// Graphics Params
func Palette_index(palette_index int32) int32
func Sprite_sheet_index(sprite_sheet_index int32) int32
func Sprite_index(sprite_index int32) int32
func Color_index(color_index int32) int32
func Flip_x(flip_x int32) int32
func Flip_y(flip_y int32) int32
func Graphics_parameters(
	palette_index int32,
	sprite_sheet_index int32,
	sprite_index int32,
	color_index int32,
	flip_x int32,
	flip_y int32,
) int32

// Draw
func Clear_screen(graphics_parameters int32)
func Set_pixel(graphics_parameters int32, x int32, y int32)
func Circle(graphics_parameters int32, x int32, y int32, radius int32)
func Circle_filled(graphics_parameters int32, x int32, y int32, radius int32)
func Rect(graphics_parameters int32, x int32, y int32, width int32, height int32)
func Rect_filled(graphics_parameters int32, x int32, y int32, width int32, height int32)
func Line(graphics_parameters int32, x0 int32, y0 int32, x1 int32, y1 int32)
func Sprite(graphics_parameters int32, transparency_mask int64, x int32, y int32)

// Text
func Console_log(text_ptr []uint8, len uint)

// Random
func Set_seed(seed int32)
func Random_int_range(min int32, max int32) int32
func Random_float() float32
func Random_float_range(min float32, max float32) float32

// Input
func Button_a_pressed(player_id int32) int32
func Button_a_released(player_id int32) int32
func Button_a_held(player_id int32) int32
func Button_b_pressed(player_id int32) int32
func Button_b_released(player_id int32) int32
func Button_b_held(player_id int32) int32
func Button_c_pressed(player_id int32) int32
func Button_c_released(player_id int32) int32
func Button_c_held(player_id int32) int32
func Button_d_pressed(player_id int32) int32
func Button_d_released(player_id int32) int32
func Button_d_held(player_id int32) int32
func Button_up_pressed(player_id int32) int32
func Button_up_released(player_id int32) int32
func Button_up_held(player_id int32) int32
func Button_down_pressed(player_id int32) int32
func Button_down_released(player_id int32) int32
func Button_down_held(player_id int32) int32
func Button_left_pressed(player_id int32) int32
func Button_left_released(player_id int32) int32
func Button_left_held(player_id int32) int32
func Button_right_pressed(player_id int32) int32
func Button_right_released(player_id int32) int32
func Button_right_held(player_id int32) int32
func Button_start_pressed(player_id int32) int32
func Button_start_released(player_id int32) int32
func Button_start_held(player_id int32) int32
func Button_select_pressed(player_id int32) int32
func Button_select_released(player_id int32) int32
func Button_select_held(player_id int32) int32
func Button_left_shoulder_pressed(player_id int32) int32
func Button_left_shoulder_released(player_id int32) int32
func Button_left_shoulder_held(player_id int32) int32
func Button_right_shoulder_pressed(player_id int32) int32
func Button_right_shoulder_released(player_id int32) int32
func Button_right_shoulder_held(player_id int32) int32
func Button_left_stick_pressed(player_id int32) int32
func Button_left_stick_released(player_id int32) int32
func Button_left_stick_held(player_id int32) int32
func Button_right_stick_pressed(player_id int32) int32
func Button_right_stick_released(player_id int32) int32
func Button_right_stick_held(player_id int32) int32
func Button_left_trigger_pressed(player_id int32) int32
func Button_left_trigger_released(player_id int32) int32
func Button_left_trigger_held(player_id int32) int32
func Button_right_trigger_pressed(player_id int32) int32
func Button_right_trigger_released(player_id int32) int32
func Button_right_trigger_held(player_id int32) int32
func Analog_left_x(player_id int32) float32
func Analog_left_y(player_id int32) float32
func Analog_right_x(player_id int32) float32
func Analog_right_y(player_id int32) float32
func Trigger_left(player_id int32) float32
func Trigger_right(player_id int32) float32
func Raw_input_state(played_id int32) int64

// Multiplayer
func Num_players() int32
func Is_local_player(player_id int32) int32
func Is_remote_player(player_id int32) int32
