
### activity/doc.go
<pre>
</pre>

### activity/activity.go
<pre>
<code>func NewActivity(dev device.Device) Activity</code>
<code>func (am Activity) StartActivity(canonicalClass string, options ...string) error</code>
<code>func (am Activity) GetFocusedActivity() (string, error)</code>
<code>func (am Activity) IsActivityFocused(name string) (bool, error)</code>
<code>func (am Activity) WaitForActivityToFocus(name string, timeout int) error</code>
</pre>

### adbutility/doc.go
<pre>
</pre>

### adbutility/adbutility.go
<pre>
<code>func GetNewAdbEndpoint(adb string) AdbEndpoint</code>
<code>func (ep AdbEndpoint) Adb(timeout int, args ...string) (string, error)</code>
<code>func (ep AdbEndpoint) GetAttachedDevices(timeout int) ([]string, error)</code>
<code>func (ep AdbEndpoint) WaitForSerials(timeout int, serials ...string) error</code>
<code>func (ep AdbEndpoint) WaitForDevices(timeout int, count int) error</code>
</pre>

### device/device.go
<pre>
<code>func NewDevice(serial string, timeout int, endPoint adbutility.AdbEndpoint) Device</code>
<code>func (dev Device) IsAvailable() (bool, error)</code>
<code>func (dev Device) Adb(command string, args ...string) (string, error)</code>
<code>func (dev Device) Shell(command string, args ...string) (string, error)</code>
<code>func (dev Device) GetProperty(key string) (string, error)</code>
<code>func (dev Device) GetAllProperties() (map[string]string, error)</code>
<code>func (dev Device) Pull(src string, dst string) (string, error)</code>
<code>func (dev Device) Push(src string, dst string) (string, error)</code>
<code>func (dev Device) WaitForAvailability(timeout int) error</code>
<code>func (dev Device) Root() error</code>
<code>func (dev Device) Reboot(restartTimeout int, bootTimeout int) error</code>
<code>func (dev Device) WaitForBootToComplete(bootTimeout int) error</code>
</pre>

### device/doc.go
<pre>
</pre>

### display/display.go
<pre>
<code>func NewDisplay(dev device.Device) Display</code>
<code>func (disp Display) GetDisplaySize() (int, int, error)</code>
<code>func (disp Display) SetDisplaySize(width int, height int) error</code>
</pre>

### doc.go
<pre>
</pre>

### geometry/geometry.go
<pre>
</pre>

### goandroid.go
<pre>
<code>func GetNewAndroidManager(timeout int, adb string) AndroidManager</code>
<code>func (am AndroidManager) GetNewAndroidDevice(serial string) Android</code>
<code>func (am AndroidManager) GetAttachedAndroidDevices() ([]Android, error)</code>
</pre>

### input/key.go
<pre>
<code>func NewKey(dev device.Device) Key</code>
<code>func (key Key) Press(code int, count int) error</code>
<code>func (key Key) PressMenu(count int) error</code>
<code>func (key Key) PressHome(count int) error</code>
<code>func (key Key) PressBack(count int) error</code>
<code>func (key Key) PressCall(count int) error</code>
<code>func (key Key) PressEndCall(count int) error</code>
<code>func (key Key) PressUp(count int) error</code>
<code>func (key Key) PressDown(count int) error</code>
<code>func (key Key) PressLeft(count int) error</code>
<code>func (key Key) PressRight(count int) error</code>
<code>func (key Key) PressVolumeUp(count int) error</code>
<code>func (key Key) PressVolumeDown(count int) error</code>
<code>func (key Key) PressPower(count int) error</code>
<code>func (key Key) PressCamera(count int) error</code>
<code>func (key Key) PressEnter(count int) error</code>
<code>func (key Key) PressDelete(count int) error</code>
</pre>

### input/input.go
<pre>
<code>func NewInputManager(dev device.Device) InputManager</code>
</pre>

### input/doc.go
<pre>
</pre>

### input/touchscreen.go
<pre>
<code>func NewTouchScreen(dev device.Device) TouchScreen</code>
<code>func (ts TouchScreen) Tap(x int, y int) error</code>
<code>func (ts TouchScreen) Swipe(x1 int, y1 int, x2 int, y2 int, delay int) error</code>
<code>func (ts TouchScreen) SwipeDown(count int) error</code>
<code>func (ts TouchScreen) SwipeUp(count int) error</code>
<code>func (ts TouchScreen) SwipeLeft(count int) error</code>
<code>func (ts TouchScreen) SwipeRight(count int) error</code>
<code>func (ts TouchScreen) RawSendEvent(dev string, eventType int, event int, value int) error</code>
<code>func (ts TouchScreen) GetTouchInputDevice() (string, error)</code>
</pre>

### input/text.go
<pre>
<code>func NewTextInput(dev device.Device) TextInput</code>
<code>func (ti TextInput) EnterText(text string)</code>
</pre>

### input/gesture.go
<pre>
<code>func (ts TouchScreen) DrawGesture(points geometry.Points, delay int) error</code>
<code>func (ts TouchScreen) DrawGestureEmulator(points geometry.Points, delay int) error</code>
<code>func (ts TouchScreen) RawMovePoint(dev string, x int, y int, id int, pressure int, size int) error</code>
<code>func (ts TouchScreen) RawMovePointEmulator(dev string, x int, y int) error</code>
</pre>

### view/textoperations.go
<pre>
<code>func (devView DeviceView) IsTextPresent(text string, index int, timeout int) error</code>
<code>func (devView DeviceView) IsMatchingTextPresnt(text string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickText(text string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickMatchingText(text string, index int, timeout int) error</code>
<code>func (devView DeviceView) GetViewForText(text string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) GetViewForMatchingText(text string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) ScrollDownToText(text string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToText(text string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollDownToMatchingText(text string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToMatchingText(text string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) GetTextForResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTextForMatchingResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTextForType(typename string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTextForMatchingType(typename string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTextForDescription(description string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTextForMatchingDescription(description string, index int, timeout int) (string, error)</code>
</pre>

### view/resourceoperations.go
<pre>
<code>func (devView DeviceView) IsResourcePresent(resource string, index int, timeout int) error</code>
<code>func (devView DeviceView) IsMatchingResourcePresnt(resource string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickResource(resource string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickMatchingResource(resource string, index int, timeout int) error</code>
<code>func (devView DeviceView) GetViewForResource(resource string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) GetViewForMatchingResource(resource string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) ScrollDownToResource(resource string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToResource(resource string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollDownToMatchingResource(resource string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToMatchingResource(resource string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) GetResourceForText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetResourceForMatchingText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetResourceForType(typename string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetResourceForMatchingType(typename string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetResourceForDescription(description string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetResourceForMatchingDescription(description string, index int, timeout int) (string, error)</code>
</pre>

### view/descriptionoperations.go
<pre>
<code>func (devView DeviceView) IsDescriptionPresent(description string, index int, timeout int) error</code>
<code>func (devView DeviceView) IsMatchingDescriptionPresnt(description string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickDescription(description string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickMatchingDescription(description string, index int, timeout int) error</code>
<code>func (devView DeviceView) GetViewForDescription(description string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) GetViewForMatchingDescription(description string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) ScrollDownToDescription(description string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToDescription(description string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollDownToMatchingDescription(description string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToMatchingDescription(description string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) GetDescriptionForText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetDescriptionForMatchingText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetDescriptionForResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetDescriptionForMatchingResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetDescriptionFortype(typename string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetDescriptionForMatchingType(typename string, index int, timeout int) (string, error)</code>
</pre>

### view/view.go
<pre>
<code>func (views Views) GetByText(text string, index int) (View, bool)</code>
<code>func (views Views) GetByMatchingText(text string, index int) (View, bool)</code>
<code>func (views Views) GetByResource(resource string, index int) (View, bool)</code>
<code>func (views Views) GetByMatchingResource(resource string, index int) (View, bool)</code>
<code>func (views Views) GetByDescription(description string, index int) (View, bool)</code>
<code>func (views Views) GetByMatchingDescription(description string, index int) (View, bool)</code>
<code>func (views Views) GetByType(typename string, index int) (View, bool)</code>
<code>func (views Views) GetByMatchingType(typename string, index int) (View, bool)</code>
</pre>

### view/deviceview.go
<pre>
<code>func NewDeviceView(dev device.Device) DeviceView</code>
<code>func (devView DeviceView) GetViewes() (Views, error)</code>
<code>func (devView DeviceView) GetHierarchy() (Hierarchy, error)</code>
</pre>

### view/typeoperations.go
<pre>
<code>func (devView DeviceView) IsTypePresent(typename string, index int, timeout int) error</code>
<code>func (devView DeviceView) IsMatchingTypePresnt(typename string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickType(typename string, index int, timeout int) error</code>
<code>func (devView DeviceView) ClickMatchingType(typename string, index int, timeout int) error</code>
<code>func (devView DeviceView) GetViewForType(typename string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) GetViewForMatchingType(typename string, index int, timeout int) (View, error)</code>
<code>func (devView DeviceView) ScrollDownToType(typename string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToType(typename string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollDownToMatchingType(typename string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) ScrollUpToMatchingType(typename string, index int, maxscroll int) error</code>
<code>func (devView DeviceView) GetTypeForText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTypeForMatchingText(text string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTypeForResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTypeForMatchingResource(resource string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTypeForDescription(description string, index int, timeout int) (string, error)</code>
<code>func (devView DeviceView) GetTypeForMatchingDescription(description string, index int, timeout int) (string, error)</code>
</pre>

### view/hierarchy.go
<pre>
<code>func (hierarchy Hierarchy) ConvertToViews() (Views, error)</code>
<code>func (nodes Nodes) ConvertToViews() (Views, error)</code>
<code>func (node Node) ConvertToView() (View, error)</code>
</pre>
